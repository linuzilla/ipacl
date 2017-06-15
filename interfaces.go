package ipacl

type IPListMgmt interface {
	Contains(ipstr string) bool
	AddEntry(ipstr ...string) error
}
