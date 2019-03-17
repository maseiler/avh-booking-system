export default {
  data: function() {
    return {
      items: [],
    };
  },
  computed:{
    itemsAlc: function(){
      return this.items.filter(function(item){
        return item.Type == 'alcoholic'
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
      this.$http.get("/getUnreservedItems").then(response => {
        var temp = response.body;
        this.items = [].concat.apply([], temp)
      });
    },
  },
  created(){
    this.$nextTick(this.getItems())
  }
};