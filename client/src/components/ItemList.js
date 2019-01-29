export default {
  data: function() {
    return {
      items: [],
    };
  },
  computed:{
    itemsAlc: function(){
      return this.items.filter(function(item){
        return item.Type == 'alcohol'
      })
    },
    itemsNonAlc: function(){
      return this.items.filter(function(item){
        return item.Type == 'non-alcoholic'
      })
    },
    itemsFood: function(){
      return this.items.filter(function(item){
        return item.Type == 'food'
      })
    },
    itemsBoat: function(){
      return this.items.filter(function(item){
        return item.Type == 'boat'
      })
    }
  },
  methods: {
    getItems: function() {
      this.$http.get("/items").then(response => {
        this.items = response.body;
      });
      console.log("getItems");
    },
  },
  created(){
    this.$nextTick(this.getItems())
  }
};