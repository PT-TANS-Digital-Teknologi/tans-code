package tanscode

// @author PT Tans Digital Teknologi
// @desc this file will help on Error Code
// @file tanscode.go
// @since 21 Feb 2023

const (

	// OK This is for everything OKAY
	OK = "0000"

	// SUCCESSCREATE This is for Success Create
	SUCCESSCREATE = "0001"

	// SUCCESSUPDATE This is for Success UPDATE
	SUCCESSUPDATE = "0002"

	// SUCCESSUDELETE This is for Success DELETE
	SUCCESSUDELETE = "0003"

	// DATAEMPTY this is for Input bad request Data Empty
	DATAEMPTY = "4001"

	// DATAREQUIRED this is for input bad request data required
	DATAREQUIRED = "4002"

	// DATADUPLICATED this is for input bad request data duplicate
	DATADUPLICATED = "4003"

	// WRONGPASSWORD this is for input bad request wrong password
	WRONGPASSWORD = "4004"

	// DATANOTFOUND this is for data not found
	DATANOTFOUND = "4005"

	// USERLOCKED this is for bad request user locked
	USERLOCKED = "4006"

	// RESTRICTEDACCESS this is for restricted access
	RESTRICTEDACCESS = "4007"

	// SHORTPHONENUMBER this is bad request phone number short
	SHORTPHONENUMBER = "4008"

	// USERNEEDTOVALIDATE this is bad request user need to validate
	USERNEEDTOVALIDATE = "4009"

	// USERALREADYACTIVE this is bad request user already active
	USERALREADYACTIVE = "4010"

	// INVALIDCREDENTIALS this is for bad request invalid credentials
	INVALIDCREDENTIALS = "4011"

	// FORBIDDEN this is FORBIDEN
	FORBIDDEN = "4012"

	// COMPANYNOTFOUND this is for bad request company not found
	COMPANYNOTFOUND = "4013"

	// USERALREADYEXPIRED this is bad request user already expired
	USERALREADYEXPIRED = "4013"

	// INTERNALDBPROBLEM this is internal DB Problem
	INTERNALDBPROBLEM = "5001"

	// INTERNALJWTPROBLEM this is JWT Problem
	INTERNALJWTPROBLEM = "5002"
)
