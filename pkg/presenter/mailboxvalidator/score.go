package mailboxvalidator

import (
	"regexp"
	"strings"
)

func CalculateScore(presenter DepPresenter) float64 {
	score, minScore := 1, 1

	if presenter.IsDomain && presenter.IsSyntax {
		score += 9
		if presenter.IsSmtp {
			score += 10
		}
		if presenter.IsVerified {
			score += 40

			if presenter.IsDisposable {
				score = 30
				if presenter.IsCatchall {
					score -= 5
				}
			} else {
				if !presenter.IsFree {
					score += 39
					if presenter.IsCatchall {
						score -= 44
					} else if presenter.IsRole {
						score -= 39
					}
				} else if presenter.IsCatchall {
					score -= 5
				}
			}
		}
	}
	if score < minScore {
		score = minScore
	}

	pos := strings.Index(presenter.EmailAddress, "@")
	if pos == -1 {
		pos = len(presenter.EmailAddress) - 1
	}
	username := presenter.EmailAddress[:pos]
	hasNumberInUserNameRE := regexp.MustCompile(`.*\d.*?`)
	hasDotInUsername := regexp.MustCompile(`.*\..*?`)

	if hasNumberInUserNameRE.MatchString(username) {
		score -= 2
	}
	if hasDotInUsername.MatchString(username) {
		score += 1
	}

	return float64(score) / 100.0
}
