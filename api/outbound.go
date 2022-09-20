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
				outbound := models.Outbound_Shadowsocks{}
				db.Add_outbound_Shadowsocks(outbound)
			case "VLESS":
				outbound := models.Outbound_VLESS{}
				db.Add_outbound_VLESS(outbound)
			case "VMess":
				outbound := models.Outbound_VMess{}
				db.Add_outbound_VMess(outbound)
			case "Trojan":
				outbound := models.Outbound_Trojan{}
				db.Add_outbound_Trojan(outbound)
			case "WireGuard":
				outbound := models.Outbound_WireGuard{}
				db.Add_outbound_WireGuard(outbound)
			case "Hysteria":
				outbound := models.Outbound_Hysteria{}
				db.Add_outbound_Hysteria(outbound)
			}
		}
	}
}

func api_outbound_mod(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.FormValue("action") == "mod" {
			switch r.FormValue("type") {
			case "Shadowsocks":
				outbound := models.Outbound_Shadowsocks{}
				db.Add_outbound_Shadowsocks(outbound)
			case "VLESS":
				outbound := models.Outbound_VLESS{}
				db.Add_outbound_VLESS(outbound)
			case "VMess":
				outbound := models.Outbound_VMess{}
				db.Add_outbound_VMess(outbound)
			case "Trojan":
				outbound := models.Outbound_Trojan{}
				db.Add_outbound_Trojan(outbound)
			case "WireGuard":
				outbound := models.Outbound_WireGuard{}
				db.Add_outbound_WireGuard(outbound)
			case "Hysteria":
				outbound := models.Outbound_Hysteria{}
				db.Add_outbound_Hysteria(outbound)
			}
		}
	}
}

func api_outbound_del(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.FormValue("action") == "del" {

			outbound := models.Del_outbound{}
			outbound.Action = r.FormValue("action")
			outbound.ID = r.FormValue("id")

			if db.Del_outbound(outbound) {
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
