package asciicoolness

import "encoding/json"

// A custom type to collect our results
// remember that anything with a capital letter will be exported.
type asciiresult struct {
	ASCIINumber  int
	ASCIICharter string
}

// a parent type to encapsulate the result type and provide a container receiver to our function
type ASCIICoolness struct {
	NumbersList []int
	ASCIIResult []asciiresult // Notice this slice is the container of our results
	Length      int           //  we can populate this however we like
}

// A method on our type that populates our results with ascii information
func (m *ASCIICoolness) populate() {
	for i := 0; i < m.Length; i++ {
		resultitem := asciiresult{ASCIINumber: i, ASCIICharter: string(i)}
		m.ASCIIResult = append(m.ASCIIResult, resultitem)
	}
}

// Marshall is our only public method an importer will see
func (m *ASCIICoolness) MarshalJSON() ([]byte, error) {
	// call to the private method to populate our values
	m.populate()
	// marshal our type to json
	out, err := json.Marshal(m.ASCIIResult)
	// simplly return the error to the caller for handling if an error occurs
	if err != nil {
		return nil, err
	}
	// return the json data
	return out, nil
}
