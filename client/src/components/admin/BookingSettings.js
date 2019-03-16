import BookingSettings from "./BookingSettings.vue"
import LastBookings from "./LastBookings.vue"

export default {
    components: {
        BookingSettings,
        LastBookings
    },
    props:{
        lastBookings: [],
        users: [],
        items: []
    },
    methods: {
    }
}