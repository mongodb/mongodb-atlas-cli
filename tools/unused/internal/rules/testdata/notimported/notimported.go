package notimported

func ExportedAndImported() {

}

func notExported() {

}

func ExportedAndNevertImported() { // want `Exported func but never imported`
	notExported()
}
