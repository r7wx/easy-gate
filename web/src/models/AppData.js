/**
 * Model of easy-gate.json.
 */
export class AppData {
    constructor(title, icon, motd, services, categories, notes, theme) {
        this.title = title || "";
        this.icon = icon || "";
        this.motd = motd || "";
        this.services = services || [];
        this.categories = categories || [];
        this.notes = notes || [];
        this.theme = theme || {
            background: "#FFFFFF",
            foreground: "#000000",
        };
        this.error = [];
    }
}