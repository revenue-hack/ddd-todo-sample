package tododm

import (
	"unicode/utf8"

	"github.com/revenue-hack/ddd-todo-sample/src/core/domain/userdm"
	"github.com/revenue-hack/ddd-todo-sample/src/core/domain/vo"
	"golang.org/x/xerrors"
)

type Todo struct {
	id            TodoID
	title         string
	status        Status
	contents      string
	asigneeUserID userdm.UserID
	referenceURL  vo.URL
}

const (
	maxTitleLength    = 100
	maxContentsLength = 1000
)

func titleValidation(title string) error {
	// マルチバイトも許可するため
	if utf8.RuneCountInString(title) > maxTitleLength {
		return xerrors.Errorf("title is less and equal than %d", maxTitleLength)
	}
	return nil
}

func contentsValidation(contents string) error {
	// マルチバイトも許可するため
	if utf8.RuneCountInString(contents) > maxContentsLength {
		return xerrors.Errorf("contents is less and equal than %d", maxContentsLength)
	}
	return nil
}

func newTodo(id TodoID, title string, status Status, contents string, asigneeUserID userdm.UserID, referenceURL vo.URL) (*Todo, error) {
	if err := titleValidation(title); err != nil {
		return nil, err
	}
	if err := contentsValidation(title); err != nil {
		return nil, err
	}

	return &Todo{
		id:            id,
		title:         title,
		status:        status,
		contents:      contents,
		asigneeUserID: asigneeUserID,
		referenceURL:  referenceURL,
	}, nil
}

func (t *Todo) ID() TodoID {
	return t.id
}
func (t *Todo) Title() string {
	return t.title
}
func (t *Todo) Status() Status {
	return t.status
}
func (t *Todo) Contents() string {
	return t.contents
}
func (t *Todo) AsigneeUserID() userdm.UserID {
	return t.asigneeUserID
}
func (t *Todo) ReferenceURL() vo.URL {
	return t.referenceURL
}
func (t *Todo) Equals(t2 *Todo) bool {
	return t.title == t2.title
}

func (t *Todo) ChangeTitle(title string) error {
	if err := titleValidation(title); err != nil {
		return err
	}
	t.title = title
	return nil
}

func (t *Todo) ChangeContents(contents string) error {
	if err := contentsValidation(contents); err != nil {
		return err
	}
	t.contents = contents
	return nil
}

func (t *Todo) ChangeStatus(s Status) {
	t.status = s
}
func (t *Todo) ChangeReferenceURL(url vo.URL) {
	t.referenceURL = url
}
