package settings

func (suitor *Suitor) TestFetchFile()  {
	file, err := suitor.loader.fetchFile()
	suitor.NoError(err)
	suitor.Equal("test.dev.json", file)
}