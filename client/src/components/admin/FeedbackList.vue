<template>
  <div>
    <div class="column">
      <table class="table is-hoverable is-striped">
        <thead>
          <tr>
            <th>ID</th>
            <th>Timestamp</th>
            <th>Name</th>
            <th>Text</th>
            <th></th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="entry in feedbackList"
            :key="entry"
            :class="[selectedEntry === entry ? 'is-selected' : '']"
          >
            <th>{{entry.ID}}</th>
            <th>{{printDateTime(entry.TimeStamp)}}</th>
            <td>
              <pre  style="white-space: pre-wrap; word-break: keep-all;">{{entry.Name}}</pre>
            </td>
            <td>
              <pre style="white-space: pre-wrap; word-break: keep-all;">{{entry.Content}}</pre>
            </td>
            <td>
              <button class="button is-link" @click="deleteEntry(entry.ID)">Delete</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  computed: {
    feedbackList() {
      return this.$store.state.feedback;
    }
  },
  data() {
    return {
      selectedEntry: {}
    };
  },
  methods: {
    deleteEntry(id) {
      this.$http
        .post("deleteFeedback", id)
        .then(() => {
          this.$responseEventBus.$emit(
            "successMessage",
            "Deleted feedback entry."
          );
          this.$store.commit("getFeedbackList");
        })
        .catch(() => {
          this.$responseEventBus.$emit(
            "failureMessage",
            "Couldn't delete feedback."
          );
        });
    }
  }
};
</script>
