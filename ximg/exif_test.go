package ximg

//func TestExif(t *testing.T) {
//	src := "../logo.png"
//	dst := path.Join(os.TempDir(), "logo.png")
//	xfile.Write(dst, xfile.Read(src))
//	err := WriteExif(dst, map[string]string{
//		"XResolution": "36",
//	})
//	assert.Nil(t, err)
//
//	exifData := ReadExif(dst)
//	assert.Less(t, 0, len(exifData))
//	assert.Equal(t, "image/png", exifData["MIMEType"])
//	assert.Equal(t, float64(36), exifData["XResolution"])
//	os.Remove(dst)
//}
