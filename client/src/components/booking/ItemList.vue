<template>
    <div>
        <p class="title">Items</p>
        <h4 class="subtitle is-4">Alcohol</h4>
        <div class="dropdown is-active">
            <div class="dropdown-trigger">
                <button class="button" is-primary v-for="item in itemsAlc" :key="item">{{ item.Name + " " + item.Size + " " + item.Unit }}</button>
            </div>
        </div>
          
        <h4 class="subtitle is-4">No Alcohol</h4>
        <div class="dropdown is-active">
            <div class="dropdown-trigger">
                <button class="button" is-primary v-for="item in itemsNonAlc" :key="item">{{ item.Name  + " " + item.Size + " " + item.Unit  }}</button>
            </div>
        </div>
          
          <h4 class="subtitle is-4">Food</h4>
          <div class="dropdown is-active">
            <div class="dropdown-trigger">
              <button class="button" is-primary v-for="item in itemsFood" :key="item">{{ item.Name }}</button>
            </div>
          </div>
          
        <h4 class="subtitle is-4">Boats</h4>
        <div class="dropdown is-active">
            <div class="dropdown-trigger">
                <button class="button" is-primary v-for="item in itemsBoat" :key="item">{{ item.Name }}</button>
            </div>
        </div>
    </div>
</template>

 <script src="./ItemList.js"></script>