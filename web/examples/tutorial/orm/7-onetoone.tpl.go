//** This file was code generated by got. DO NOT EDIT. ***

package orm

import (
	"bytes"
	"context"

	"github.com/goradd/goradd/web/examples/gen/goradd/model"
	"github.com/goradd/goradd/web/examples/gen/goradd/model/node"
)

func (ctrl *OneOnePanel) DrawTemplate(ctx context.Context, buf *bytes.Buffer) (err error) {

	buf.WriteString(`
<h1>One-to-One Relationships</h1>
<p>
By creating a unique index on a foreign key, you will create a one-to-one relationship.
One-to-one relationships link two records with each record being an extension of the other.
You can use this to model subclasses, extend a record with optional values, or to improve search results, allowing the database
to retrieve only the main record on an initial search, and then the extended record if more detail is desired of
a particular record.
</p>
<p>
In the example below, we load a login, and then examine which person the login belongs to.
`)
	login := model.LoadLogin(ctx, "2", node.Login().Person())

	buf.WriteString(`</p>
<p>
    Person for Login `)

	buf.WriteString(login.Username())

	buf.WriteString(`: `)

	buf.WriteString(login.Person().FirstName())

	buf.WriteString(` `)

	buf.WriteString(login.Person().LastName())

	buf.WriteString(` <br>
</p>
<p>
Here, we traverse the relationship in the other direction, loading the person first, and then getting the login.
`)
	person := model.LoadPerson(ctx, "3", node.Person().Login())

	buf.WriteString(`</p>
<p>
    Login for `)

	buf.WriteString(person.FirstName())

	buf.WriteString(` `)

	buf.WriteString(person.LastName())

	buf.WriteString(`: `)

	buf.WriteString(person.Login().Username())

	buf.WriteString(`  <br>
</p>

<h2>Creating One-to-One Linked Records</h2>
<p>
In a similar fashion to how one-to-many relationships work, you can create a link between two records by saving one,
getting its id, and then setting the foreign key in the other record to that id. However, its easier to use the Set*
functions for the objects themselves and call Save on the parent object.
</p>
`)
	newPerson := model.NewPerson()
	newPerson.SetFirstName("Hu")
	newPerson.SetLastName("Man")

	newLogin := model.NewLogin()
	newLogin.SetUsername("human")

	newPerson.SetLogin(newLogin)
	newPerson.Save(ctx)

	buf.WriteString(`<p>
    New person `)

	buf.WriteString(newPerson.FirstName())

	buf.WriteString(` `)

	buf.WriteString(newPerson.LastName())

	buf.WriteString(` has been given a Login ID of `)

	buf.WriteString(newPerson.Login().ID())

	buf.WriteString(`
</p>
`)
	// Delete records created above
	newPerson.Delete(ctx) // newLogin will automatically get deleted because its foreign key constraint is set to CASCADE on Delete

	buf.WriteString(`
`)

	return
}
