<template>

  <div class="columns p-2 pt-4">
    <!-- add param -->
    <b-sidebar
      type="is-light"
      :fullheight="true"
      :overlay="false"
      :open.sync="showAddParam"
      :can-cancel="['escape', 'x']"
    >
      <div class="p-2 pt-4">
        <b-field label="Operator" label-position="on-border">
          <b-autocomplete
            v-model="param.operator"
            placeholder="e.g. >="
            :data="filteredOperators"
            :open-on-focus="true"
            @select="option => (selected = option)"
          >
          </b-autocomplete>
        </b-field>

        <b-field label="Key" label-position="on-border">
          <input
            class="input"
            v-model="param.key"
            type="text"
            placeholder="Field name"
          />
        </b-field>

        <b-field
          label="Value"
          label-position="on-border"
          v-if="isListValue(param.operator)"
        >
          <compos :vmodstring="param.value.list" phtext="Value"></compos>
          <!-- <b-taginput
            v-model="param.value.list"
            autocomplete
            :allow-new="true"
            placeholder="Value"
            icon="label"
          > -->
          <!-- </b-taginput> -->
        </b-field>
        <b-field label="Value" label-position="on-border" v-else>
          <input
            class="input"
            v-model="param.value.item"
            type="text"
            placeholder="Value"
          />
        </b-field>
        <button
          class="button is-small is-fullwidth is-outlined is-success"
          @click="saveParam()"
        >
          Add
        </button>
      </div>
    </b-sidebar>
    <!-- add param -->
    <div class="column page-menu">
      <!-- date -->
      <b-field label="Start time" label-position="on-border">
        <b-datetimepicker
          placeholder="Click to select..."
          :max-datetime="maxDatetime"
          :timepicker="{ enableSeconds: true }"
          editable
          v-model="form.startTime"
        ></b-datetimepicker>
      </b-field>

      <b-field label="End time" label-position="on-border">
        <b-datetimepicker
          placeholder="Click to select..."
          :timepicker="{ enableSeconds: true }"
          editable
          v-model="form.endTime"
        ></b-datetimepicker>
      </b-field>
      <!-- // date -->

      <!-- level -->
      <b-field label="Level =" label-position="on-border">
        <compos :vmodstring="form.level" :datar="levels" phtext="Level"></compos>
        <!-- <b-taginput
          v-model="form.level"
          :data="levels"
          autocomplete
          :allow-new="true"
          :open-on-focus="true"
          placeholder="Level"
          @typing="getLevelList"
          icon="label"
        > -->
        <!-- </b-taginput> -->
      </b-field>
      <!-- // level -->

      <!-- namespace -->
      <b-field label="Namespace =" label-position="on-border">
        <compos :vmodstring="form.namespace" :datar="namespaces" phtext="Namespace"></compos>
        <!-- <b-taginput
          v-model="form.namespace"
          :data="namespaces"
          autocomplete
          :allow-new="true"
          :open-on-focus="true"
          placeholder="Namespace"
          @typing="getNamespaceList"
          icon="label"
        > -->
        <!-- </b-taginput> -->
      </b-field>
      <!-- // namespace -->

      <!-- source -->
      <b-field label="Source" label-position="on-border">
         <compos :vmodstring="form.source" :datar="sources" phtext="Source"></compos>
        <!-- <b-taginput
          v-model="form.source"
          :data="sources"
          autocomplete
          :allow-new="true"
          :open-on-focus="true"
          placeholder="Source"
          @typing="getSourceList"
          icon="label"
        > -->
        <!-- </b-taginput> -->
      </b-field>
      <!-- // source -->
      <!-- traceID -->
      <b-field label="Trace ID" label-position="on-border">
         <compos :vmodstring="form.traceID" phtext="Trace ID"></compos>
        <!-- <b-taginput
          v-model="form.traceID"
          autocomplete
          :allow-new="true"
          placeholder="Trace ID"
          icon="label"
        >
        </b-taginput> -->
      </b-field>
      <!-- // traceID -->

      <!-- params -->
      <b-field
        v-for="(param, i) in params"
        :label="`${param.key} ${param.operator}`"
        :key="`param_${i}`"
        label-position="on-border"
      >
        <!-- <b-taginput
          v-if="isListValue(param.operator)"
          v-model="param.value.list"
          autocomplete
          :allow-new="true"
          placeholder="Value"
          icon="label"
          icon-right="close-circle"
          icon-right-clickable
          @icon-right-click="removeParam(i)"
        >
        </b-taginput> -->
        <compos v-if="isListValue(param.operator)"
          :vmodstring="param.value.list"
          phtext="Value"
          icon-right="close-circle"
          icon-right-clickable
          @icon-right-click="removeParam(i)"
        >
        </compos>
        <b-input
          v-else
          :placeholder="param.key"
          v-model="param.value.item"
          type="text"
          icon-right="close-circle"
          icon-right-clickable
          @icon-right-click="removeParam(i)"
        >
        </b-input>
      </b-field>
      <!-- // params -->

      <template v-if="showAdditionalParam">
        <!-- host -->
        <b-field label="Host" label-position="on-border">
          <compos :vmodstring="form.host" :datar="hosts" phtext="Host"></compos>
          <!-- <b-taginput
            v-model="form.host"
            :data="hosts"
            autocomplete
            :allow-new="true"
            :open-on-focus="true"
            placeholder="Host"
            @typing="getHostList"
            icon="label"
          >
          </b-taginput> -->
        </b-field>
        <!-- // host -->

        <!-- Build commit -->
        <b-field label="Build commit" label-position="on-border">
          <compos :vmodstring="form.buildCommit" phtext="Build commit"></compos>
          <!-- <b-taginput
            v-model="form.buildCommit"
            autocomplete
            :allow-new="true"
            placeholder="Build commit"
            icon="label"
          > -->
          <!-- </b-taginput> -->
        </b-field>
        <!-- // Build commit -->

        <!-- Config Hash -->
        <b-field label="Config hash" label-position="on-border">
          <compos :vmodstring="form.configHash" phtext="Config hash"></compos>
          <!-- <b-taginput
            v-model="form.configHash"
            autocomplete
            :allow-new="true"
            placeholder="Config hash"
            icon="label"
          > -->
          <!-- </b-taginput> -->
        </b-field>
        <!-- // Config Hash -->
      </template>

      <div class="buttons is-centered">
        <button
          class="button is-small is-outlined"
          @click="showAdditionalParam = !showAdditionalParam"
        >
          <b-icon
            :icon="showAdditionalParam ? 'eye-off' : 'eye'"
            size="is-small"
          >
          </b-icon>
          <span>other</span>
        </button>
        <button
          class="button is-small is-outlined"
          @click="showAddParam = true"
        >
          <b-icon icon="plus" size="is-small"> </b-icon>
          <span>param</span>
        </button>
      </div>
      <b-button class="button is-primary is-fullwidth" @click="search"
        >Search</b-button
      >
    </div>

    <div class="column">
      <b-field label="Search" label-position="on-border">
        <b-input
          placeholder="Search..."
          type="search"
          icon="magnify"
          icon-clickable
          class="w100"
          v-model="form.message"
        ></b-input>
        <p class="control">
          <b-button class="button is-primary">Search</b-button>
        </p>
      </b-field>

      <p v-for="(m, i) in messages" :key="i">{{ JSON.stringify(m) }}</p>

      <b-skeleton
        size="is-large"
        :active="loading"
        :count="20"
        v-if="messages.length === 0"
      ></b-skeleton>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Param, Form, ParamValue } from '../../types/view';
import compos from '../../components/ComponentExample.vue';

export default Vue.extend({
  components: {
    compos,
  },
  data() {
    return {
      loading: true,
      form: {
        startTime: new Date(new Date().getTime() - 1000 * 60),
        endTime: null,
        // namespace: [] as string[],
        // source: [] as string[],
        // traceID: [] as string[],
        // host: [] as string[],
        // level: [] as string[],
        // buildCommit: '',
        // configHash: '',
        message: '',
      } as Form,
      params: [] as Param[],
      param: {
        operator: '',
        type: '',
        key: '',
        value: {
          item: '',
          list: [] as string[],
        } as ParamValue,
      } as Param,
      maxDatetime: new Date(),
      sources: [],
      hosts: [],
      namespaces: [],
      levels: [],
      operators: ['=', '!=', '>', '<', '>=', '<=', 'LIKE', 'NOT LIKE'],
      showAddParam: false,
      showAdditionalParam: false,
      messages: [],
    };
  },
  computed: {
    filteredOperators(): string[] {
      return this.operators.filter(
        (option: string) => option.toLowerCase().indexOf(this.param.operator.toLowerCase()) >= 0,
      );
    },
  },
  methods: {
    // getSourceList(val: string): void {
    //   console.log(val);
    // },
    // getHostList(val: string): void {
    //   console.log(val);
    // },
    // getNamespaceList(val: string): void {
    //   console.log(val);
    // },
    // getLevelList(val: string): void {
    //   console.log(val);
    // },
    saveParam(): void {
      this.showAddParam = false;

      this.params.push({
        type: this.param.type,
        key: this.param.key,
        value: this.param.value,
        operator: this.param.operator,
      } as Param);

      this.param = {
        operator: '',
        type: '',
        key: '',
        value: {
          item: '',
          list: [] as string[],
        } as ParamValue,
      } as Param;
    },
    removeParam(idx: number): void {
      this.params = this.params.filter((v, i) => i !== idx);
    },
    isListValue(operator: string): boolean {
      return ['=', '!=', 'LIKE', 'NOT LIKE'].includes(operator);
    },
    search(): void {
      const params = [
        {
          type: 'column',
          key: 'time',
          operator: '>=',
          value: {
            item: parseInt(
              (this.form.startTime.getTime() / 1000).toString(),
              10,
            ).toString(),
          } as ParamValue,
        },
      ] as Param[];

      if (this.form.endTime !== null) {
        params.push({
          type: 'column',
          key: 'time',
          operator: '<=',
          value: { item: this.form.endTime } as ParamValue,
        });
      }

      if (this.form.message !== '') {
        params.push({
          type: 'column',
          key: 'message',
          operator: 'LIKE',
          value: { item: this.form.message } as ParamValue,
        });
      }

      [
        { key: 'namespace', value: this.form.namespace },
        { key: 'level', value: this.form.level },
        { key: 'source', value: this.form.source },
        { key: 'trace_id', value: this.form.traceID },
        { key: 'host', value: this.form.host },
        { key: 'build_commit', value: this.form.buildCommit },
        { key: 'config_hash', value: this.form.configHash },
      ].forEach((h) => {
        if (h.value.length > 0) {
          params.push({
            type: 'column',
            key: h.key,
            operator: '=',
            value: { list: h.value } as ParamValue,
          });
        }
      });

      this.params.forEach((param) => {
        params.push({
          type: 'key',
          key: param.key,
          operator: param.operator,
          value: param.value,
        });
      });

      Vue.axios
        .post('/api/v1/entry/list', { params, limit: 1000 })
        .then((response) => {
          this.messages = response.data.data;

          console.log(response.data);
        })
        .catch((e) => {
          console.error(e);
        });
    },
  },
});
</script>

<style lang="scss" scoped>
.p-2 {
  padding: 0.5rem;
}
.pt-4 {
  padding-top: 1rem;
}

.w100 {
  width: 100%;
}

.page {
  &-menu {
    max-width: 210px;
    min-width: 150px;
  }
}
</style>
