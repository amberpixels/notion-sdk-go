package notion

// UserID stands for ID of User object.
// User is an Object, so UserID is just an alias for Object
type UserID = ObjectID

// UserType stands for Type of User object.
type UserType string

// nolint:revive
const (
	UserTypePerson UserType = "person"
	UserTypeBot    UserType = "bot"
)

// User is a Notion object that represents a user.
type User struct {
	AtomObject
	AtomID

	Type      UserType `json:"type,omitempty"`
	Name      string   `json:"name,omitempty"`
	AvatarURL string   `json:"avatar_url,omitempty"`

	Person *Person `json:"person,omitempty"`
	Bot    *Bot    `json:"bot,omitempty"`
}

// Users is a list of Users
type Users []*User

// IsPerson returns treu if the User is a Person.
func (u *User) IsPerson() bool { return u.Type == UserTypePerson }

// IsBot returns true if the User is a Bot.
func (u *User) IsBot() bool { return u.Type == UserTypeBot }

// NewPersonUser returns a new User(Person) with the given ID and email.
func NewPersonUser(id UserID, email string) *User {
	return &User{
		AtomObject: AtomObject{Object: ObjectTypeUser},
		AtomID:     AtomID{ID: id},
		Type:       UserTypePerson,
		Person:     &Person{Email: email},
	}
}

// NewBotUser returns a new User(Bot) with the given ID and Bot instance.
func NewBotUser(id UserID, bot *Bot) *User {
	return &User{
		AtomObject: AtomObject{Object: ObjectTypeUser},
		AtomID:     AtomID{ID: id},
		Type:       UserTypeBot,
		Bot:        bot,
	}
}

// Person represents a Person.
type Person struct {
	Email string `json:"email"`
}

// Bot represents a Bot.
type Bot struct {
	Owner         Owner  `json:"owner"`
	WorkspaceName string `json:"workspace_name"`
}

// Owner represents a Bot's Owner.
type Owner struct {
	Type      string `json:"type"`
	Workspace bool   `json:"workspace"`
}
