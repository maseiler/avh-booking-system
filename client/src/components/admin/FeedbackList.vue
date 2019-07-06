<template>
  <div>
    <div class="column">
      <table class="table is-hoverable is-striped">
        <thead>
          <tr>
            <th>ID</th>
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
            <td>
              <pre>{{entry.Name}}</pre>
            </td>
            <td>
              <pre>{{entry.Content}}</pre>
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
        .then(function(response) {
          this.$responseEventBus.$emit(
            "successMessage",
            "Deleted feedback entry."
          );
          this.$store.commit("getFeedbackList");
        })
        .catch(function(response) {
          this.$responseEventBus.$emit(
            "failureMessage",
            "Couldn't delete feedback."
          );
        });
    }
  }
};
</script>
