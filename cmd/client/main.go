// ============================================================
// FILE: cmd/client/main.go
// ROLE: Demo client. Submits two C++ programs to the judge server
//       and prints the verdict for each.
//
// FLOW:
//   submit(label, code)
//     [1] JSON-encode  { "code": "<C++ source>" }
//     [2] POST http://localhost:8080/submit
//     [3] Read + decode JSON response
//     [4] Print result:
//           ✔  AC  ──▶  print "Accepted" + actual output
//           ✘  WA  ──▶  print "Wrong Answer" + actual output + diff
//
// TWO TEST CASES:
//   correctCode  — reads 5 nums, prints them in order  →  expect AC
//   wrongCode    — reads 5 nums, prints them REVERSED  →  expect WA
// ============================================================
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// correctCode reads 5 numbers and prints them back — should get AC
const correctCode = `#include <iostream>
using namespace std;
int main() {
    int a, b, c, d, e;
    cin >> a >> b >> c >> d >> e;
    cout << a << " " << b << " " << c << " " << d << " " << e << endl;
    return 0;
}
`

// wrongCode reads 5 numbers but prints them in reverse — should get WA
const wrongCode = `#include <iostream>
using namespace std;
int main() {
    int a, b, c, d, e;
    cin >> a >> b >> c >> d >> e;
    cout << e << " " << d << " " << c << " " << b << " " << a << endl;
    return 0;
}
`

type SubmissionRequest struct {
	Code string `json:"code"`
}

type SubmissionResponse struct {
	Verdict string `json:"verdict"`
	Diff    string `json:"diff,omitempty"`
	Output  string `json:"output,omitempty"`
	Error   string `json:"error,omitempty"`
}

func submit(label, code string) {
	fmt.Printf("\n========== Submission: %s ==========\n", label)

	body, _ := json.Marshal(SubmissionRequest{Code: code})
	resp, err := http.Post("http://localhost:8080/submit", "application/json", bytes.NewReader(body))
	if err != nil {
		fmt.Fprintf(os.Stderr, "HTTP error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)

	var res SubmissionResponse
	if err := json.Unmarshal(raw, &res); err != nil {
		fmt.Printf("Raw response: %s\n", raw)
		return
	}

	if res.Error != "" {
		fmt.Printf("❌ Server error: %s\n", res.Error)
		return
	}

	if res.Verdict == "AC" {
		fmt.Printf("✅ Verdict: AC (Accepted)\n")
		fmt.Printf("   Output : %s\n", res.Output)
	} else {
		fmt.Printf("❌ Verdict: WA (Wrong Answer)\n")
		fmt.Printf("   Output : %s\n", res.Output)
		fmt.Printf("   Diff   :\n%s\n", res.Diff)
	}
}

func main() {
	submit("Correct solution (echo)", correctCode)
	submit("Wrong solution (reversed)", wrongCode)
}
