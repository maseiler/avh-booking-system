<template>
  <div v-if="item !== 'undefined' && Object.keys(item).length !== 0">
    <article class="message is-link">
      <div class="message-header">
        <p>{{ item.Name }}</p>
      </div>
      <div class="message-body">
        <div class="columns">
          <div class="column has-text-left">
            <h6 class="subtitle is-6">
              <div>
                <font-awesome-icon icon="fingerprint" />
                &nbsp;{{ item.ID }}
              </div>
              <div>
                <font-awesome-icon icon="font" />
                &nbsp;{{ item.Name }}
              </div>
              <div>
                <font-awesome-icon icon="euro-sign" />
                &nbsp;&nbsp;{{ $n(item.Price, "currency", "de-DE") }}
              </div>
            </h6>
          </div>
          <div class="column has-text-left">
            <h6 class="subtitle is-6">
              <div>
                <font-awesome-icon icon="expand-arrows-alt" />
                &nbsp;&nbsp;{{ item.Size }}
              </div>
              <div>
                <font-awesome-icon icon="balance-scale" />
                &nbsp;{{ item.Unit }}
              </div>
              <div>
                <font-awesome-icon icon="info-circle" />
                &nbsp;&nbsp;{{ item.Type }}
              </div>
            </h6>
          </div>
          <div class="column has-text-left">
            <h6 class="subtitle is-6">
              <div>
                <font-awesome-icon v-if="item.Enabled" icon="check-circle"/>
                <font-awesome-icon v-if="!item.Enabled" icon="times"/>
                &nbsp;&nbsp;{{ item.Enabled }}
              </div>
            </h6>
          </div>
        </div>
      </div>
    </article>
  </div>
</template>

<script>
export default {
  computed: {
    item() {
      return this.$store.state.selectedSingleItem;
    },
  },
};
</script>