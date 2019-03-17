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
  },
  props: {
    items: []
  },
  data: function () {
    return {
      showAddItemForm: false,
      showModifyItemForm: false,
      showDeleteItemForm: false,
      search: '',
      searchResults: [],
      selectedItem: {}
    };
  },
  methods: {
    searchItems: function () {
      if (this.search != '') {
        var tmpSearch = this.search.toLowerCase()
        this.searchResults = this.items.filter(item => (item['Name'].toLowerCase().includes(tmpSearch)) | (item['Type'].toLowerCase().includes(tmpSearch)))
      } else {
        this.searchResults = []
      }
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