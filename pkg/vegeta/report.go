package vegeta

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"vegeta-server/models"

	vegeta "github.com/tsenart/vegeta/lib"
)

// Format defines a type for the format query param
type Format string

const (
	// JSONFormat typedef for query param "json"
	JSONFormat Format = "json"
	// TextFormat typedef for query param "text"
	TextFormat Format = "text"
	// HistogramFormat typedef for query param "histogram"
	//HistogramFormat Format = "histogram"
	// BinaryFormat typedef for query param "binary"
	BinaryFormat Format = "binary"
)

// CreateReportFromReader takes in an io.Reader with the vegeta gob, encoded result and
// returns the decoded result as a byte array
func CreateReportFromReader(reader io.Reader, id string, format Format) ([]byte, error) {
	dec := vegeta.DecoderFor(reader)

	m := vegeta.Metrics{}

	var report vegeta.Report = &m

decode:
	for {
		var r vegeta.Result
		err := dec.Decode(&r)
		if err != nil {
			if err == io.EOF {
				break decode
			}
			return nil, err
		}

		report.Add(&r)
	}

	rc := report.(vegeta.Closer)
	rc.Close()

	var rep vegeta.Reporter

	switch format {
	case JSONFormat:
		// Create a new reporter with the metrics
		rep = vegeta.NewJSONReporter(&m)
	case TextFormat:
		rep = vegeta.NewTextReporter(&m)
	// TODO: Figure out how to provide historgram report
	//case HistogramFormat:
	//      var hist vegeta.Histogram
	//      if err := hist.Buckets.UnmarshalText([]byte(typ[4:])); err != nil {
	//              return err
	//      }
	default:
		return nil, fmt.Errorf("format %s not supported", format)
	}

	var b []byte
	buf := bytes.NewBuffer(b)
	err := rep.Report(buf)
	if err != nil {
		return nil, err
	}

	if format == JSONFormat {
		// Add ID to the report struct
		var jsonReportResponse models.JSONReportResponse
		err = json.Unmarshal(buf.Bytes(), &jsonReportResponse)
		if err != nil {
			return nil, err
		}
		jsonReportResponse.ID = id
		return json.Marshal(jsonReportResponse)
	} else if format == TextFormat {
		return addID(buf, id), nil
	}

	return buf.Bytes(), nil
}

func addID(report *bytes.Buffer, id string) []byte {
	return append([]byte(fmt.Sprintf("ID %s\n", id)), report.Bytes()...)
}
