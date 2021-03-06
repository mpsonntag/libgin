package libgin

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepoPathToUUID(t *testing.T) {
	mkey := "fabee/efish_locking"
	mval := "6953bbf0087ba444b2d549b759de4a06"

	retval := RepoPathToUUID(mkey)
	if retval != mval {
		t.Fatalf("Expected %s but got %s", mval, retval)
	}

	mkey = "ppfeiffer/cooperative_channels_in_biological_neurons_via_dynamic_clamp"
	mval = "08853efd775dc85d9636bae2043c8f2c"
	retval = RepoPathToUUID(mkey)
	if retval != mval {
		t.Fatalf("Expected %s but got %s", mval, retval)
	}
}

func TestGetType(t *testing.T) {
	var res = DOIRegInfo{}

	check := res.GetType()
	if check != "Dataset" {
		t.Fatalf("Expected 'Dataset' but got %q", check)
	}

	val := "Datapaper"
	res.ResourceType = val
	check = res.GetType()
	if check != val {
		t.Fatalf("Expected %q but got %q", val, check)
	}
}

func TestAuthor(t *testing.T) {
	lname := "ln"
	fname := "fn"
	aff := "af"
	id := "id"

	var validate string
	var auth Author
	check := func(label string) {
		rend := auth.RenderAuthor()
		if validate != rend {
			t.Fatalf("%s: Expected '%s' but got '%s'", label, validate, rend)
		}
	}

	// Test RenderAuthor; make all expected combinations explicit
	// No omit test
	validate = fmt.Sprintf("%s, %s; %s; %s", lname, fname, aff, id)

	auth = Author{
		FirstName:   fname,
		LastName:    lname,
		Affiliation: aff,
		ID:          id,
	}
	check("No omit")

	// Omit ID test
	validate = fmt.Sprintf("%s, %s; %s", lname, fname, aff)

	auth = Author{
		FirstName:   fname,
		LastName:    lname,
		Affiliation: aff,
	}
	check("Omit ID")

	// Omit affiliation test
	validate = fmt.Sprintf("%s, %s; %s", lname, fname, id)

	auth = Author{
		FirstName: fname,
		LastName:  lname,
		ID:        id,
	}
	check("Omit affiliation")

	// Omit ID and affiliation test
	validate = fmt.Sprintf("%s, %s", lname, fname)

	auth = Author{
		FirstName: fname,
		LastName:  lname,
	}
	check("Omit ID/affiliation")
}

func TestGetURL(t *testing.T) {
	// check empty or malformed ID
	var ref = Reference{}

	checkstr := ref.GetURL()
	if checkstr != "" {
		t.Fatalf("Expected empty URL on empty ID but got %q", checkstr)
	}
	ref.ID = "IamNotOK"
	checkstr = ref.GetURL()
	if checkstr != "" {
		t.Fatalf("Expected empty URL on invalid ID but got %q", checkstr)
	}

	// check empty on empty or unknown prefix
	ref.ID = ":value"
	checkstr = ref.GetURL()
	if checkstr != "" {
		t.Fatalf("Expected empty URL on missing prefix but got %q", checkstr)
	}
	ref.ID = "you:dontKnowMe"
	checkstr = ref.GetURL()
	if checkstr != "" {
		t.Fatalf("Expected empty URL on unknown ID prefix but got %q", checkstr)
	}

	// check url prefix behavior
	idval := "reference"
	ref.ID = fmt.Sprintf("uRl:%s", idval)
	checkstr = ref.GetURL()
	if checkstr != idval {
		t.Fatalf("Expected %q on URL prefix but got %q", idval, checkstr)
	}

	// check doi prefix behavior
	ref.ID = fmt.Sprintf("DOI:%s", idval)
	resstr := fmt.Sprintf("doi.org/%s", idval)
	checkstr = ref.GetURL()
	if !strings.Contains(checkstr, resstr) {
		t.Fatalf("Got unexpected DOI URL string %q", checkstr)
	}

	// check arxiv prefix behavior
	ref.ID = fmt.Sprintf("arXiv:%s", idval)
	checkstr = ref.GetURL()
	resstr = fmt.Sprintf("arxiv.org/abs/%s", idval)
	if !strings.Contains(checkstr, resstr) {
		t.Fatalf("Got unexpected arxiv URL string %q", checkstr)
	}

	// check pmid prefix behavior
	ref.ID = fmt.Sprintf("pmID:%s", idval)
	checkstr = ref.GetURL()
	resstr = fmt.Sprintf("www.ncbi.nlm.nih.gov/pubmed/%s", idval)
	if !strings.Contains(checkstr, resstr) {
		t.Fatalf("Got unexpected pmid URL string %q", checkstr)
	}
}

func TestIsRegisteredDOI(t *testing.T) {
	invalid := "idonotexist"
	valid := "10.12751/g-node.5b08du"

	// check false on non-existing DOI
	ok := IsRegisteredDOI(invalid)
	if ok {
		t.Fatal("Expected check to fail on invalid DOI")
	}

	ok = IsRegisteredDOI(valid)
	if !ok {
		t.Fatal("Expected check to succeed on valid DOI")
	}
}
