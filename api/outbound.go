package api

import (
	"Sing-boxA/db"
	"Sing-boxA/models"
	"net/http"
)

func api_outbound_add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.FormValue("action") == "add" {
			switch r.FormValue("type") {
			case "Shadowsocks":
				addOutbound := models.Outbound_Shadowsocks{}
				db.Add_Outbound_Shadowsocks(addOutbound)
			case "VLESS":
				addOutbound := models.Outbound_VLESS{}
				db.Add_Outbound_VLESS(addOutbound)
			case "VMess":
				addOutbound := models.Outbound_VMess{}
				db.Add_Outbound_VMess(addOutbound)
			case "Trojan":
				addOutbound := models.Outbound_Trojan{}
				db.Add_Outbound_Trojan(addOutbound)
			case "WireGuard":
				addOutbound := models.Outbound_WireGuard{}
				db.Add_Outbound_WireGuard(addOutbound)
			case "Hysteria":
				addOutbound := models.Outbound_Hysteria{}
				db.Add_Outbound_Hysteria(addOutbound)
			}
		}
	}
}

func api_outbound_mod(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.FormValue("action") == "mod" {
			switch r.FormValue("type") {
			case "Shadowsocks":
				modOutbound := models.Outbound_Shadowsocks{}
				db.Add_Outbound_Shadowsocks(modOutbound)
			case "VLESS":
				modOutbound := models.Outbound_VLESS{}
				db.Add_Outbound_VLESS(modOutbound)
			case "VMess":
				modOutbound := models.Outbound_VMess{}
				db.Add_Outbound_VMess(modOutbound)
			case "Trojan":
				modOutbound := models.Outbound_Trojan{}
				db.Add_Outbound_Trojan(modOutbound)
			case "WireGuard":
				modOutbound := models.Outbound_WireGuard{}
				db.Add_Outbound_WireGuard(modOutbound)
			case "Hysteria":
				modOutbound := models.Outbound_Hysteria{}
				db.Add_Outbound_Hysteria(modOutbound)
			}
		}
	}
}

func api_outbound_del(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.FormValue("action") == "del" {

			delOutbound := models.Del_Outbound{}
			delOutbound.Action = r.FormValue("action")
			delOutbound.ID = r.FormValue("id")

			if db.Del_Outbound(delOutbound) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte(`{"info": "success"}`))
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte(`{"info": "fail"}`))
			}
		}
	}
}
