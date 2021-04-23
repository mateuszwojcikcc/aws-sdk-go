// +build go1.7

package lookoutmetrics

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/aws/mateuszwojcikcc/aws"
	"github.com/aws/mateuszwojcikcc/awstesting/unit"
)

func TestClientContentType(t *testing.T) {
	sess := unit.Session.Copy()

	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			contentType := r.Header.Get("Content-Type")
			if e, a := contentType, "application/x-amz-json-1.1"; !strings.EqualFold(e, a) {
				t.Errorf("expect %v content-type, got %v", e, a)
			}
		},
	))
	defer server.Close()

	client := New(sess, &aws.Config{Endpoint: &server.URL})
	_, err := client.ActivateAnomalyDetector(&ActivateAnomalyDetectorInput{
		AnomalyDetectorArn: aws.String("foo"),
	})
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}
}
