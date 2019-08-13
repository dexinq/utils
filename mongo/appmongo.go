package mongo

// app mongo data action
func (m Mongodb) InsertAppEntity(data interface{}) error {
	return nil
}

func (m Mongodb) UpdateAppEntity(data interface{}) error {
	return nil
}

func (m Mongodb) RemoveAppEntity(appName string) error {
	return nil
}

// task mongo data action
func (m Mongodb)InsertTaskEntity(data interface{}) error {
	return nil
}

func (m Mongodb)RemoveTaskEntity(taskName string) error {
	return nil
}

func (m Mongodb)UpdateTaskStatus(taskName string, status int32) error {
	return nil
}

func (m Mongodb)UpdateTaskEndTime(taskName string, endTime int64) error {
	return nil
}


