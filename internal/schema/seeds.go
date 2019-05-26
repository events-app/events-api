package schema

// seeds is a string constant containing all of the queries needed to get the
// db seeded to a useful state for development.
//
// Using a constant in a .go file is an easy way to ensure the queries are part
// of the compiled executable and avoids pathing issues with the working
// directory. It has the downside that it lacks syntax highlighting and may be
// harder to read for some cases compared to using .sql files. You may also
// consider a combined approach using a tool like packr or go-bindata.
//
// Note that database servers besides PostgreSQL may not support running
// multiple queries as part of the same execution so this single large constant
// may need to be broken up.

const seeds = `
INSERT INTO card (name, text, created) VALUES 
	('main', '# The New Event

	The New Event is the best event ever.
	You should definitelly attend!
	
	+ Register at [The New Event](http://thenewevent.com/).
	+ Come
	+ Have fun
	
	We are waiting for you!', '2019-05-26 00:00:02.000001+00'),
	('secured', 'You are allowed to see this.', '2019-05-26 00:00:02.000001+00'),
	('other', 'Other content', '2019-05-26 00:00:02.000001+00'),
	('other2','Other content 2', '2019-05-26 00:00:02.000001+00')
	ON CONFLICT DO NOTHING;

INSERT INTO menu (name, card_id, created) VALUES
	('Main menu', 1, '2019-05-26 00:00:02.000001+00'),
	('Secured', 2, '2019-05-26 00:00:02.000001+00'),
	('Other menu', 3, '2019-05-26 00:00:02.000001+00')
	ON CONFLICT DO NOTHING;
`

// TODO add seeds for user and file
