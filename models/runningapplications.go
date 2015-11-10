package models

type RunningApplication struct {
  Id       					int64
  Name 						  string					`sql:"size:255; not null; unique;"`
  ApplicationID			int64 					
  Owner				      UserApp			    `sql:"not null;"`
  AccessUrl          string         `sql:"size:255; not null;"` 
  	Username				string					`sql:"-"`
  	Message					string 					`sql:"-"`	//These dont get put in the database
}

type UserApp struct {
	ID 				int64
  UserID    int64	
	RunningApplicationID	int64 	`sql:index`
}