import ItemSearch from "./ItemSearch.vue"

export default {
    components: {
        ItemSearch
    },
    props: {
        items: []
    },
    data: function () {
        return {
            search: '',
            searchResults: [],
            selectedItem: {}
        };
    },
    methods: {
        searchItems: function(){
            if(this.search != ''){
                var tmpSearch = this.search.toLowerCase()
                this.searchResults = this.items.filter(item => (item['Name'].toLowerCase().includes(tmpSearch)))
            } else {
                this.searchResults = []
            }
        },
        displayItem: function (item) {
            if (item.Type == "alcoholic" || item.Type == "non-alcoholic") {
                return item.Name + " " + item.Size + " " + item.Unit
            } else if (item.Type == "boat" || item.Type == "food") {
                return item.Name
            } else {
                return "???"
            }
        },
        selectItem: function(item){
            this.selectedItem = item
            this.$emit('selectItem', item)
        }
    }
};