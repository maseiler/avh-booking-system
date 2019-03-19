import Test from "./Test.vue"

export default {
    components: {
        Test
    }, props: {
        value: String,
        endpoint: String,
    },
    computed: {
        displayValue: {
            get() {
                return this.value
            },
            set(value) {
                this.$emit('input', value)
            },
        },
    },
};