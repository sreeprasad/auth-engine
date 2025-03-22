package auth

func (s *Statement) Matches(r *Request) bool {
	principalMatch := false
	for _, p := range s.Principal {
		if WildcardMatch(p, r.Principal) {
			principalMatch = true
			break
		}
	}
	if !principalMatch {
		return false
	}

	actionMatch := false
	for _, a := range s.Action {
		if WildcardMatch(a, r.Action) {
			actionMatch = true
			break
		}
	}
	if !actionMatch {
		return false
	}

	resourceMatch := false
	for _, res := range s.Resource {
		if WildcardMatch(res, r.Resource) {
			resourceMatch = true
			break
		}
	}

	return resourceMatch

}

func MatchesDeny(policies []Policy, r *Request) bool {
	for _, p := range policies {
		for _, s := range p.Statements {
			if s.Effect == Deny && s.Matches(r) {
				return true
			}
		}
	}
	return false
}

func MatchesAllow(policies []Policy, r *Request) bool {
	for _, p := range policies {
		for _, s := range p.Statements {
			if s.Effect == Allow && s.Matches(r) {
				return true
			}
		}
	}
	return false
}

func Evaluate(policies []Policy, r *Request) Answer {
	if MatchesDeny(policies, r) {
		return ExplicitDeny
	} else if MatchesAllow(policies, r) {
		return AllowAccess
	} else {
		return ImplicitDeny
	}
}
