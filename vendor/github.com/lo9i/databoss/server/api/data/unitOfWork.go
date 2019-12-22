package data

type UnitOfWork interface {
	Database() Database
	Begin()
	Commit()
}

type UnitOfWorkImpl struct {
	database Database
}

func NewUnitOfWork(database Database) UnitOfWork {
	return &UnitOfWorkImpl{ database:database }
}

func (u UnitOfWorkImpl) Database() Database {
	return u.database
}

func (u UnitOfWorkImpl) Begin()  {
	u.database.Begin()
}

func (u UnitOfWorkImpl) Commit()  {
	u.database.Commit()
}