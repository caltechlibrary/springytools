// libguides_test.go implements tests for springytools LibGuides related
// data structures.
package springytools

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"testing"
)

func codingSequence(t *testing.T, srcName string, destName string, obj interface{}) error {
	var (
		src []byte
		err error
	)
	src, err = ioutil.ReadFile(srcName)
	if err != nil {
		t.Errorf("read %q: %s", srcName, err)
		return err
	}
	err = xml.Unmarshal(src, &obj)
	if err != nil {
		t.Errorf("xml Unmarshal %q: %s", srcName, err)
		return err
	}
	src, err = json.MarshalIndent(obj, "", "    ")
	if err != nil {
		t.Errorf("json Marshal %q: %s", destName, err)
		return err
	}
	err = ioutil.WriteFile(destName, src, 0777)
	if err != nil {
		t.Errorf("writing %q: %s", destName, err)
		return err
	}
	return nil
}

func expectedInt(t *testing.T, expected int, got int) {
	if expected != got {
		t.Errorf("expected %d, got %d", expected, got)
	}
}

func expectedString(t *testing.T, expected string, got string) {
	if expected != got {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

func TestCustomer(t *testing.T) {
	var customer *Customer
	customer = new(Customer)
	if err := codingSequence(t, "testinput/customer.xml", "testout/customer.json", customer); err != nil {
		t.FailNow()
	}
	expectedInt(t, 64, customer.Id)
	expectedString(t, "Academic Institution", customer.Type)
	expectedString(t, "https://library.example.edu/", customer.Url)
	expectedString(t, "Anytown", customer.City)
	expectedString(t, "Euforia", customer.State)
	expectedString(t, "United Places of North America", customer.Country)
	expectedString(t, "America/Los_Angeles", customer.TimeZone)
	expectedString(t, "2014-02-13 00:24:29", customer.Created)
	expectedString(t, "2020-02-04 19:59:10", customer.Updated)
}

func TestSite(t *testing.T) {
	var site *Site
	site = new(Site)
	if err := codingSequence(t, "testinput/site.xml", "testout/site.json", site); err != nil {
		t.FailNow()
	}
	expectedInt(t, 64, site.Id)
	expectedString(t, "LibGuides", site.Type)
	expectedString(t, "LibGuides", site.Name)
	expectedString(t, "libguides.example.edu", site.Domain)
	expectedString(t, "libapps@library.example.edu", site.Admin)
	expectedString(t, "2014-02-13 00:24:29", site.Created)
	expectedString(t, "2020-07-21 22:00:18", site.Updated)
}

func TestAccount(t *testing.T) {
	var account *Account
	account = new(Account)
	if err := codingSequence(t, "testinput/account-1.xml", "testout/account-1.json", account); err != nil {
		t.FailNow()
	}
	expectedInt(t, 1, account.Id)
	expectedString(t, "shrimps@engineering.example.edu", account.Email)
	expectedString(t, "Crusty", account.FirstName)
	expectedString(t, "Anthropod", account.LastName)
	expectedString(t, "A Watery Engineer", account.Title)
	expectedString(t, "Barnicle Bob", account.Nickname)
	expectedString(t, `Somewhere in the food chain | MC 0-07 | Anytown, Euforia 0000001 | 111-222-3333 | www.library.example.edu `, account.Signature)
	expectedString(t, "https://caltechlibrary.github.io/", account.Website)
	expectedString(t, "", account.Image)
	expectedString(t, "", account.Skype)
	expectedString(t, "", account.Address)
	expectedString(t, "2019-09-09 22:37:09", account.Created)
	expectedString(t, "2020-06-29 15:24:34", account.Updated)
}

func TestAccounts(t *testing.T) {
	var accounts []*Account
	if err := codingSequence(t, "testinput/accounts.xml", "testout/accounts.json", &accounts); err != nil {
		t.FailNow()
	}
}

func TestGroups(t *testing.T) {
	groups := []*Group{}
	if err := codingSequence(t, "testinput/groups.xml", "testout/groups.json", &groups); err != nil {
		t.FailNow()
	}
}

func TestSubjects(t *testing.T) {
	subjects := []*Subject{}
	if err := codingSequence(t, "testinput/subjects.xml", "testout/subjects.json", &subjects); err != nil {
		t.FailNow()
	}
}

func TestTags(t *testing.T) {
	tags := []*Tag{}
	if err := codingSequence(t, "testinput/tags.xml", "testout/tags.json", &tags); err != nil {
		t.FailNow()
	}
}

func TestVendors(t *testing.T) {
	vendors := []*Vendor{}
	if err := codingSequence(t, "testinput/vendors.xml", "testout/vendors.json", &vendors); err != nil {
		t.FailNow()
	}
}

func TestPage(t *testing.T) {
	var page *Page
	page = new(Page)
	if err := codingSequence(t, "testinput/page-3502869.xml", "testout/page-3502869.json", page); err != nil {
		t.FailNow()
	}
}

func TestLibGuides(t *testing.T) {
	libguides := LibGuides{}
	if err := codingSequence(t, "testinput/LibGuides_export_XXXXX.xml", "testout/LibGuide_export.json", &libguides); err != nil {
		t.FailNow()
	}
}
