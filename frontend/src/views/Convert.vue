<template>
  <div class="col-sm">
    <div v-if="loading" class="text-center">
      <img src="../assets/loading.gif">
    </div>
    <b-form @submit="onSubmit" v-if="!loading">
      <b-form-group
        id="convertWidth"
        label="Enter the desired width:"
        label-for="exampleInput1"
        description="Size is described in pixels"
      >
        <b-form-input
          id="convertWidthInput"
          type="number"
          v-model="form.width"
          required
          placeholder="Enter email"
        ></b-form-input>
      </b-form-group>
      <b-form-group
        id="convertHeight"
        label="Enter the desired height (in px):"
        label-for="convertHeightInput"
        description="Size is described in pixels"
      >
        <b-form-input
          id="convertHeightInput"
          type="number"
          v-model="form.height"
          required
          placeholder="Enter name"
        ></b-form-input>
      </b-form-group>
      <b-form-group
        id="convertFile"
        label="Select the file you want to convert:"
        label-for="cconvertFileInput"
      >
        <b-form-file
          id="cconvertFileInput"
          v-model="form.file"
          placeholder="Choose a file..."
          accept="image/*"
        ></b-form-file>
      </b-form-group>
      <div class="text-center">
        <b-button type="submit" variant="primary">Submit</b-button>
      </div>
    </b-form>
  </div>
</template>
<script>
import JobService from '@/services/JobService';

export default {
  data() {
    return {
      loading: false,
      form: {
        width: 200,
        height: 200,
        preserveRatio: false,
        file: null
      }
    };
  },
  methods: {
    onSubmit(evt) {
      this.loading = true;
      evt.preventDefault();
      const data = new FormData();
      data.append('width', this.form.width);
      data.append('file', this.form.file);
      data.append('height', this.form.height);
      JobService.create(data)
        .then(() => {
          this.loading = false;
          this.$notify({
            group: 'all',
            title: 'Image sent!',
            text: 'Our engines are now working on your image!'
          });
        })
        .catch(e => {
          alert('An error occured. Please open the dev tools to debug it.');
          console.error(e);
        });
    }
  }
};
</script>
