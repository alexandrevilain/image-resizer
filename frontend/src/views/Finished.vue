<template>
  <div class="col-sm">
    <b-table striped responsive hover :items="items" :fields="fields">      
      <template slot="originalUrl" slot-scope="data">
        <img :src=data.value class="img-fluid"/>
      </template>
      <template slot="resultUrl" slot-scope="data">
        <img :src=data.value class="img-fluid"/>
      </template>
    </b-table>
  </div>
</template>
<script>
import ImagesService from '@/services/ImagesService';

export default {
  data() {
    return {
      fields: [
        {
          key: 'desiredWidth',
          label: 'Desired Width'
        },
        {
          key: 'desiredHeight',
          label: 'Desired Height'
        },
        {
          key: 'originalUrl',
          label: 'Original'
        },
        {
          key: 'resultUrl',
          label: 'Result'
        },
        {
          key: 'finishedat',
          formatter: value => {
            return new Date(value).toLocaleString();
          }
        }
      ],
      items: []
    };
  },
  beforeMount() {
    ImagesService.getAll().then(({ data }) => {
      this.items = data;
    });
  }
};
</script>
