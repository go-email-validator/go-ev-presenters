package adapter

import (
	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/go-email-validator/go-ev-presenters/pkg/presentation/as_email_verifier"
)

func ResultAdaption(presentation as_email_verifier.DepPresentation) emailverifier.Result {
	var smtp *emailverifier.SMTP

	if presentation.SMTP == nil {
		smtp = &emailverifier.SMTP{
			HostExists:  presentation.SMTP.HostExists,
			FullInbox:   presentation.SMTP.FullInbox,
			CatchAll:    presentation.SMTP.CatchAll,
			Deliverable: presentation.SMTP.Deliverable,
			Disabled:    presentation.SMTP.Disabled,
		}
	}

	var gravatar *emailverifier.Gravatar
	if presentation.Gravatar == nil {
		gravatar = &emailverifier.Gravatar{
			HasGravatar: presentation.Gravatar.HasGravatar,
			GravatarUrl: presentation.Gravatar.GravatarUrl,
		}
	}

	return emailverifier.Result{
		Email:       presentation.Email,
		Disposable:  presentation.Disposable,
		Reachable:   presentation.Reachable.String(),
		RoleAccount: presentation.RoleAccount,
		Free:        presentation.Free,
		Syntax: &emailverifier.Syntax{
			Username: presentation.Syntax.Username,
			Domain:   presentation.Syntax.Domain,
			Valid:    presentation.Syntax.Valid,
		},
		HasMxRecords: presentation.HasMxRecords,
		SMTP:         smtp,
		Gravatar:     gravatar,
		Suggestion:   presentation.Suggestion,
	}
}
