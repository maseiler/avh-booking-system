<template>
  <transition name="modal">
    <div class="modal-mask">
      <div class="modal-wrapper">
        <div class="modal-container">
          <div class="modal-header">
            <slot name="header">
              <h1>Delete Item</h1>
              <article v-if="Object.keys(item).length === 0" class="message is-danger">
                <div class="message-header">
                  <div class="field is-grouped">
                    <p class="icon is-small is-left">
                      <font-awesome-icon icon="exclamation" size="lg"/>
                    </p>
                    <p>Select item first.</p>
                  </div>
                </div>
              </article>
              <div v-else>
                <p>Do you really want to delete</p>
                <p>{{item.Name}} {{item.Price}}€ {{item.Size}}{{item.Unit}} ({{item.Type}})?</p>
              </div>
            </slot>
          </div>

          <div class="modal-footer">
            <slot name="footer">
              <div class="field is-grouped">
                <div class="control" v-if="Object.keys(item).length !== 0">
                  <button class="button is-link" @click="deleteItem">Delete</button>
                </div>
                <div class="control">
                  <button class="button" @click="cancelSubmission">Cancel</button>
                </div>
              </div>
            </slot>
          </div>
        </div>
      </div>
    </div>
  </transition>
</template>

<script src="./DeleteItemForm.js"></script>
<style lang="scss">
@import "../assets/modal.css";
</style>