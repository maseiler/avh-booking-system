import BookingSettings from "./BookingSettings.vue"
import LastBookings from "./LastBookings.vue"
import ModifyBookEntryForm from "./ModifyBookEntryForm.vue"

export default {
    components: {
        BookingSettings,
        LastBookings,
        ModifyBookEntryForm
    },
    props: {
        lastBookings: [],
        users: [],
        items: [],
    },
    data: function () {
        return {
            selectedEntry: {},
            showModifyBookEntryForm: false
        }
    },
    methods: {
        selectEntry: function (entry) {
            this.selectedEntry = entry
        }
    }
}