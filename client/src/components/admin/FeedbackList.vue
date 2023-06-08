<template>
  <div class="tile is-ancestor is-narrow">
    <div class="tile is-parent is-vertical">
      <article class="tile is-child">
        <div class="column">
          <div
            class="box"
            style="
              height: calc(100vh - 3.5rem);
              overflow-x: hidden;
              overflow-y: auto;
            "
          >
            <table class="table is-hoverable is-striped">
              <thead>
                <tr>
                  <th>{{$t('generic.id')}}</th>
                  <th>{{$t('booking.payment.timestamp')}}</th>
                  <th>{{$t('generic.name')}}</th>
                  <th>{{$t('generic.text')}}</th>
                  <th></th>
                </tr>
              </thead>
              <tbody>
                <tr
                  v-for="entry in feedbackList"
                  :key="entry"
                  :class="[selectedEntry === entry ? 'is-selected' : '']"
                >
                  <th>{{ entry.ID }}</th>
                  <th>{{ printDateTime(entry.TimeStamp) }}</th>
                  <td>
                    <pre style="white-space: pre-wrap; word-break: keep-all">{{
                      entry.Name
                    }}</pre>
                  </td>
                  <td>
                    <pre style="white-space: pre-wrap; word-break: keep-all">{{
                      entry.Content
                    }}</pre>
                  </td>
                  <td>
                    <button
                      class="button is-link"
                      @click="deleteEntry(entry.ID)"
                    >
                    {{$t('generic.delete')}}
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </article>
    </div>
  </div>
</template>

<script>
export default {
  computed: {
    feedbackList() {
      return this.$store.state.feedback;
    },
  },
  data() {
    return {
      selectedEntry: {},
    };
  },
  methods: {
    deleteEntry(id) {
      this.$http
        .post("deleteFeedback", id)
        .then(() => {
          this.$responseEventBus.$emit(
            "successMessage",
            this.$t('messages.success.deleteFeedback')
          );
          this.$store.commit("getFeedbackList");
        })
        .catch(() => {
          this.$responseEventBus.$emit(
            "failureMessage",
            this.$t('messages.failure.feedbackDelete')
          );
        });
    },
  },
};
</script>
