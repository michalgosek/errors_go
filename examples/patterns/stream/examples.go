package stream

func ReadStreamExample1() error {
	bb1 := []byte{'H', 'e', 'l', 'o', '\n'}
	bb2 := []byte{'W', 'o', 'r', 'l', 'd', '\n'}
	bb3 := []byte{'G', 'o', '\n'}
	bb4 := []byte{'i', 's', '\n'}
	bb5 := []byte{'c', 'o', 'o', 'l', '\n'}
	var sr Stream

	err := sr.Read(bb1...)
	if err != nil {
		return err
	}
	err = sr.Read(bb2...)
	if err != nil {
		return err
	}
	err = sr.Read(bb3...)
	if err != nil {
		return err
	}
	err = sr.Read(bb4...)
	if err != nil {
		return err
	}
	err = sr.Read(bb5...)
	if err != nil {
		return err
	}
	return nil
}

func ReadStreamExample2() error {
	var err error
	var sr Stream

	read := func(bb ...byte) {
		if err != nil {
			return
		}
		err = sr.Read(bb...)
	}

	bb1 := []byte{'H', 'e', 'l', 'o', '\n'}
	bb2 := []byte{'W', 'o', 'r', 'l', 'd', '\n'}
	bb3 := []byte{'G', 'o', '\n'}
	bb4 := []byte{}
	bb5 := []byte{'c', 'o', 'o', 'l', '\n'}

	read(bb1...)
	read(bb2...)
	read(bb3...)
	read(bb4...)
	read(bb5...)
	if err != nil {
		return err
	}

	return nil
}

func ReadStreamExample3() error {
	bb1 := []byte{'H', 'e', 'l', 'o', '\n'}
	bb2 := []byte{}
	bb3 := []byte{'G', 'o', '\n'}
	bb4 := []byte{'i', 's', '\n'}
	bb5 := []byte{'c', 'o', 'o', 'l', '\n'}

	sr := NewErrStreamReader()

	sr.Read(bb1...)
	sr.Read(bb2...)
	sr.Read(bb3...)
	sr.Read(bb4...)
	sr.Read(bb5...)
	if sr.Error() != nil {
		return sr.Error()
	}
	return nil
}
