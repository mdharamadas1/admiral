/*
 * Copyright (c) 2017 VMware, Inc. All Rights Reserved.
 *
 * This product is licensed to you under the Apache License, Version 2.0 (the "License").
 * You may not use this product except in compliance with the License.
 *
 * This product may include a number of subcomponents with separate copyright notices
 * and license terms. Your use of these subcomponents is subject to the terms and
 * conditions of the subcomponent's license, as noted in the LICENSE file.
 */

export default Vue.component('multicolumn-editor-group', {
  template: `
    <form-group
      :class="class">
      <form-label
        :required="required">
        {{label}}
      </form-label>
      <multicolumn-editor
        :disabled="disabled"
        :headers="headers"
        :value="value"
        @change="onChange">
        <slot></slot>
      </multicolumn-editor>
    </form-group>
  `,
  props: {
    class: {
      required: false,
      type: String
    },
    disabled: {
      default: false,
      required: false,
      type: Boolean
    },
    headers: {
      required: false,
      type: Array
    },
    label: {
      required: true,
      type: String
    },
    required: {
      default: false,
      required: false,
      type: Boolean
    },
    value: {
      required: true,
      type: Array
    }
  },
  methods: {
    onChange(value) {
      this.value = value;
      this.$emit('change', this.value);
    }
  }
});

