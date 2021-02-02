// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

<template>
    <div class="container">
        <p class="container__item name" :title="itemData.name">{{ itemData.name }}</p>
        <p class="container__item description">{{ itemData.description }}</p>
        <p class="container__item member-count">{{ itemData.memberCount }}</p>
        <p class="container__item date">{{ itemData.createdDate() }}</p>
    </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';

import { Project } from '@/types/projects';

@Component
export default class ProjectsListItem extends Vue {
    @Prop({default: () => new Project('123', 'name', 'desc')})
    private readonly itemData: Project;
}
</script>

<style scoped lang="scss">
    .container {
        padding: 20px 40px;
        outline: none;
        display: flex;
        background: #fff;
        margin-bottom: 1px;
        width: calc(100% - 80px);
        cursor: pointer;

        &:hover {
            background-color: rgba(242, 244, 247, 0.6);
        }

        &__item {
            width: 25%;
            font-family: 'font_regular', sans-serif;
            font-size: 16px;
            margin: 0;
        }
    }

    .name {
        width: 25%;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
        font-family: 'font_medium', sans-serif;
    }

    .description {
        margin-left: 14px;
    }

    .member-count {
        position: relative;
        left: 20px;
    }

    .date {
        position: relative;
        left: 20px;
    }
</style>
