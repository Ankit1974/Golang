package howtomakegetrequestpractice


myUrl = "jnerjnjrnjre/lee.//reo..//kpokre"

response,err:= http.Get(myUrl)

if err!= null{
	panic(err)
}

defer response.Body.Close()

answer, err := io.ReadAll(response.Body)

if err!= null{
	panic(err)
}

