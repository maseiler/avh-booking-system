import LastBookings from "./LastBookings.vue"

export default {
    components: {
        LastBookings
    },
    props: {
        lastBookings: [],
        users: [],
        items: []
    },
    data: function () {
        return {
            selectedEntry: {}
        }
    },
    methods: {
        printDateTime: function (dateTime) {
            var d = new Date(Date.parse(dateTime))
            return ("0" + d.getDate()).slice(-2) + "." + ("0" + (d.getMonth() + 1)).slice(-2) + "." +
                d.getFullYear() + " " + ("0" + d.getHours()).slice(-2) + ":" + ("0" + d.getMinutes()).slice(-2) + ":" + ("0" + d.getSeconds()).slice(-2);
        },
        getUser: function (id) {
            return this.users.find(u => { return u.UserID == id })
        },
        displayName: function (user) {
            if (user.BierName != '') {
                return user.BierName
            } else {
                if (user.LastName != '') {
                    return user.FirstName + ' ' + user.LastName[0] + '.'
                }
                return user.FirstName
            }
        },
        getItem: function (id) {
            return this.items.find(i => { return i.ItemID == id })
        },
        displayItem: function (item) {
            if (item.Type == 'boat' || item.Type == 'food') {
                return item.Name
            } else {
                return item.Name + ' ' + item.Size + item.Unit
            }
        }
    }
}
//TODO: pass users and items as props