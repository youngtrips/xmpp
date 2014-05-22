package xmpp

import (
	"fmt"
)

const (
	xmlnsMuc      = "<x xmlns='http://jabber.org/protocol/muc'/>"
	xmlnsMucOwner = "http://jabber.org/protocol/muc#owner"
	xmlnsMucAdmin = "http://jabber.org/protocol/muc#admin"
)

func (c *Conn) CreateRoom(jid string) {
}

func (c *Conn) JoinMUC(jid string, nickname string) error {
	_, err := fmt.Fprintf(c.out, "<presence from='%s' to='%s/%s'>%s</presence>", c.jid, jid, nickname, xmlnsMuc)
	return err
}

func (c *Conn) DestroyRoom(jid string) error {
	_, err := fmt.Fprintf(c.out, "<iq from='%s' id='%x' to='%s' type='set'><query xmlns='%s'><destroy jid='%s'></destroy></query></iq>",
		c.jid, c.getCookie(), jid, xmlnsMucAdmin, jid)
	return err
}

const (
	AffiliationNone    int = 1
	AffiliationOutcast int = 2
	AffiliationMember  int = 3
	AffiliationOwner   int = 4
	AffiliationAdmin   int = 5
	AffiliationInvalid int = 6
	RoleNone           int = 1
	RoleVisitor        int = 2
	RoleParticipant    int = 3
	RoleModerator      int = 4
	RoleInvalid        int = 5
)

func (c *Conn) SetRole(roomJid, jid string, role int) error {
	roleName := ""
	switch role {
	case RoleNone:
		roleName = "none"
	case RoleVisitor:
		roleName = "visitor"
	case RoleParticipant:
		roleName = "participant"
	case RoleModerator:
		roleName = "moderator"
	case RoleInvalid:
		roleName = "invalid"
	}
	cookie := c.getCookie()
	_, err := fmt.Fprintf(c.out, "<iq from='%s' id='%x' to='%s' type='set'><query xmlns='%s'><item jid='%s' role='%s'/></query></iq>",
		c.jid, cookie, roomJid, xmlnsMucAdmin, jid, roleName)
	return err
}

func (c *Conn) SetAffiliation(roomJid, jid string, affiliation int) error {
	affiliationName := ""
	switch affiliation {
	case AffiliationNone:
		affiliationName = "none"
	case AffiliationOutcast:
		affiliationName = "outcast"
	case AffiliationMember:
		affiliationName = "member"
	case AffiliationOwner:
		affiliationName = "owner"
	case AffiliationAdmin:
		affiliationName = "admin"
	case AffiliationInvalid:
		affiliationName = "invalid"
	}
	cookie := c.getCookie()
	_, err := fmt.Fprintf(c.out, "<iq from='%s' id='%x' to='%s' type='set'><query xmlns='%s'><item affiliation='%s' jid='%s'/></query></iq>",
		c.jid, cookie, roomJid, xmlnsMucAdmin, affiliationName, jid)
	return err
}
