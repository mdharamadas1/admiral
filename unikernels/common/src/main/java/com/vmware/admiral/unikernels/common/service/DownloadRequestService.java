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

package com.vmware.admiral.unikernels.common.service;

import java.util.ArrayList;

import com.vmware.admiral.unikernels.common.service.UnikernelCreationTaskService.UnikernelCreationTaskServiceState;
import com.vmware.xenon.common.Operation;
import com.vmware.xenon.common.StatelessService;
import com.vmware.xenon.common.TaskState;
import com.vmware.xenon.common.TaskState.TaskStage;

public class DownloadRequestService extends StatelessService {

    public static final String SELF_LINK = UnikernelManagementURIParts.DOWNLOAD;
    private static ArrayList<String> downloadListURI = new ArrayList<>();

    private String OSvCompilationHostURL = "";

    @Override
    public void handlePost(Operation post) {
        String downloadURI = post.getBody(String.class);

        downloadListURI.add(OSvCompilationHostURL + downloadURI);

        UnikernelCreationTaskServiceState state = new UnikernelCreationTaskServiceState();
        state.taskInfo = new TaskState();
        state.taskInfo.stage = TaskStage.FINISHED;

        Operation request = Operation.createPatch(post.getReferer())
                .setReferer(getSelfLink())
                .setBody(state)
                .setCompletion((o, e) -> {
                    if (e != null) {
                        post.fail(e);
                    } else {
                        post.complete();
                    }
                });

        sendRequest(request);
    }

    /*
     * The method is responsible for updating the host URL, which is sent from the
     * UnikernelCreationTaskService after provisioning
     */
    @Override
    public void handlePut(Operation put) {
        OSvCompilationHostURL = put.getBody(String.class);
        put.complete();
    }

    @Override
    public void handleGet(Operation get) {
        get.setBody(downloadListURI);
        get.complete();
    }
}