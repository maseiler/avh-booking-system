import AddItemForm from "../booking/AddItemForm"
import ModifyItemForm from "./ModifyItemForm.vue"
import DeleteItemForm from "./DeleteItemForm.vue"
import ItemInfo from "./ItemInfo.vue"

export default {
  components: {
    AddItemForm,
    ModifyItemForm,
    DeleteItemForm,
    ItemInfo
  }, data: function () {
    return {
      allItems: [],
      showAddItemForm: false,
      showModifyItemForm: false,
      showDeleteItemForm: false,
      search: '',
      searchResults: [],
      selectedItem: {}
    };
  },
  methods: {
    getItems: function () {
      this.$http.get("/getItems").then(response => {
        var temp = response.body
        this.allItems = [].concat.apply([], temp)
      });
    }, searchItems: function () {
      if (this.search != '') {
        var tmpSearch = this.search.toLowerCase()
        this.searchResults = this.allItems.filter(item => (item['Name'].toLowerCase().includes(tmpSearch)) | (item['Type'].toLowerCase().includes(tmpSearch)))
      } else {
        this.searchResults = []
      }
    }, displayItem: function (item) {
      if (item.Type == 'boat' || item.Type == 'food') {
        return item.Name
      } else {
        return item.Name + ' ' + item.Size + item.Unit
      }
    }
  }, created() {
    this.$nextTick(this.getItems())
  }
}