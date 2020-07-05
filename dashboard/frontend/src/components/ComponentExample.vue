<template>

    <b-taginput
          v-model="vmodstring"
          :data=datar
          autocomplete
          :allow-new="allownew"
          :open-on-focus="opnfocus"
          :placeholder= "phtext"
          @typing="getLogList"
          icon="label"
        >
    </b-taginput>
</template>

<script lang="ts">
import { ParamValue, Form, Param } from '../types/view';

export default {
  props: {
    phtext: String,
    datar: {
      type: Array,
      default: () => [],
    },
    vmodstring: {
      type: Array,
      default: () => [],

    },
    opnfocus: {
      type: Boolean,
      default: true,
    },
    allownew: {
      type: Boolean,
      default: true,
    },
  },
  data() {
    return {
      loading: true,
      form: {
        // startTime: new Date(new Date().getTime() - 1000 * 60),
        // endTime: null,
        namespace: [] as string[],
        source: [] as string[],
        traceID: [] as string[],
        host: [] as string[],
        level: [] as string[],
        buildCommit: '',
        configHash: '',
        // message: '',
      } as Form,
      //   params: [] as Param[],
    //   param: {
    //     // operator: '',
    //     // type: '',
    //     // key: '',
    //     value: {
    //       //  item: '',
    //       list: [] as string[],
    //     } as ParamValue,
    //   } as Param,
    //   maxDatetime: new Date(),
    //   sources: [],
    //   hosts: [],
    //   namespaces: [],
    //   levels: [],
    //   operators: ['=', '!=', '>', '<', '>=', '<=', 'LIKE', 'NOT LIKE'],
    //   showAddParam: false,
    //   showAdditionalParam: false,
    //   messages: [],
    };
  },

  //   computed: {
  //     filteredOperators(): string[] {
  //       return this.operators.filter(
  //         (option: string) => option.toLowerCase().indexOf(this.param.operator.toLowerCase()) >= 0,
  //       );
  //     },
  //   },
  methods: {
    // вместо многих сделаем одну функцию которая будет создаваться для кажного компонента отдельно
    getLogList(val: string): void {
      console.log(val);
    },

  },

};
</script>

<style scoped>
</style>
