# coding: utf-8

"""
    Bacalhau API

    This page is the reference of the Bacalhau REST API. Project docs are available at https://docs.bacalhau.org/. Find more information about Bacalhau at https://github.com/bacalhau-project/bacalhau.  # noqa: E501

    OpenAPI spec version: 0.3.18.post4
    Contact: team@bacalhau.org
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


from __future__ import absolute_import

import re  # noqa: F401

# python 2 and python 3 compatibility library
import six

from bacalhau_apiclient.api_client import ApiClient


class JobApi(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    Ref: https://github.com/swagger-api/swagger-codegen
    """

    def __init__(self, api_client=None):
        if api_client is None:
            api_client = ApiClient()
        self.api_client = api_client

    def events(self, events_request, **kwargs):  # noqa: E501
        """Returns the events related to the job-id passed in the body payload. Useful for troubleshooting.  # noqa: E501

        Events (e.g. Created, Bid, BidAccepted, ..., ResultsAccepted, ResultsPublished) are useful to track the progress of a job.   # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True
        >>> thread = api.events(events_request, async_req=True)
        >>> result = thread.get()

        :param async_req bool
        :param EventsRequest events_request: Request must specify a `client_id`. To retrieve your `client_id`, you can do the following: (1) submit a dummy job to Bacalhau (or use one you created before), (2) run `bacalhau describe <job-id>` and fetch the `ClientID` field. (required)
        :return: EventsResponse
                 If the method is called asynchronously,
                 returns the request thread.
        """
        kwargs['_return_http_data_only'] = True
        if kwargs.get('async_req'):
            return self.events_with_http_info(events_request, **kwargs)  # noqa: E501
        else:
            (data) = self.events_with_http_info(events_request, **kwargs)  # noqa: E501
            return data

    def events_with_http_info(self, events_request, **kwargs):  # noqa: E501
        """Returns the events related to the job-id passed in the body payload. Useful for troubleshooting.  # noqa: E501

        Events (e.g. Created, Bid, BidAccepted, ..., ResultsAccepted, ResultsPublished) are useful to track the progress of a job.   # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True
        >>> thread = api.events_with_http_info(events_request, async_req=True)
        >>> result = thread.get()

        :param async_req bool
        :param EventsRequest events_request: Request must specify a `client_id`. To retrieve your `client_id`, you can do the following: (1) submit a dummy job to Bacalhau (or use one you created before), (2) run `bacalhau describe <job-id>` and fetch the `ClientID` field. (required)
        :return: EventsResponse
                 If the method is called asynchronously,
                 returns the request thread.
        """

        all_params = ['events_request']  # noqa: E501
        all_params.append('async_req')
        all_params.append('_return_http_data_only')
        all_params.append('_preload_content')
        all_params.append('_request_timeout')

        params = locals()
        for key, val in six.iteritems(params['kwargs']):
            if key not in all_params:
                raise TypeError(
                    "Got an unexpected keyword argument '%s'"
                    " to method events" % key
                )
            params[key] = val
        del params['kwargs']
        # verify the required parameter 'events_request' is set
        if self.api_client.client_side_validation and ('events_request' not in params or
                                                       params['events_request'] is None):  # noqa: E501
            raise ValueError("Missing the required parameter `events_request` when calling `events`")  # noqa: E501

        collection_formats = {}

        path_params = {}

        query_params = []

        header_params = {}

        form_params = []
        local_var_files = {}

        body_params = None
        if 'events_request' in params:
            body_params = params['events_request']
        # HTTP header `Accept`
        header_params['Accept'] = self.api_client.select_header_accept(
            ['application/json'])  # noqa: E501

        # HTTP header `Content-Type`
        header_params['Content-Type'] = self.api_client.select_header_content_type(  # noqa: E501
            ['application/json'])  # noqa: E501

        # Authentication setting
        auth_settings = []  # noqa: E501

        return self.api_client.call_api(
            '/requester/events', 'POST',
            path_params,
            query_params,
            header_params,
            body=body_params,
            post_params=form_params,
            files=local_var_files,
            response_type='EventsResponse',  # noqa: E501
            auth_settings=auth_settings,
            async_req=params.get('async_req'),
            _return_http_data_only=params.get('_return_http_data_only'),
            _preload_content=params.get('_preload_content', True),
            _request_timeout=params.get('_request_timeout'),
            collection_formats=collection_formats)

    def list(self, list_request, **kwargs):  # noqa: E501
        """Simply lists jobs.  # noqa: E501

        Returns the first (sorted) #`max_jobs` jobs that belong to the `client_id` passed in the body payload (by default). If `return_all` is set to true, it returns all jobs on the Bacalhau network.  If `id` is set, it returns only the job with that ID.  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True
        >>> thread = api.list(list_request, async_req=True)
        >>> result = thread.get()

        :param async_req bool
        :param ListRequest list_request: Set `return_all` to `true` to return all jobs on the network (may degrade performance, use with care!). (required)
        :return: ListResponse
                 If the method is called asynchronously,
                 returns the request thread.
        """
        kwargs['_return_http_data_only'] = True
        if kwargs.get('async_req'):
            return self.list_with_http_info(list_request, **kwargs)  # noqa: E501
        else:
            (data) = self.list_with_http_info(list_request, **kwargs)  # noqa: E501
            return data

    def list_with_http_info(self, list_request, **kwargs):  # noqa: E501
        """Simply lists jobs.  # noqa: E501

        Returns the first (sorted) #`max_jobs` jobs that belong to the `client_id` passed in the body payload (by default). If `return_all` is set to true, it returns all jobs on the Bacalhau network.  If `id` is set, it returns only the job with that ID.  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True
        >>> thread = api.list_with_http_info(list_request, async_req=True)
        >>> result = thread.get()

        :param async_req bool
        :param ListRequest list_request: Set `return_all` to `true` to return all jobs on the network (may degrade performance, use with care!). (required)
        :return: ListResponse
                 If the method is called asynchronously,
                 returns the request thread.
        """

        all_params = ['list_request']  # noqa: E501
        all_params.append('async_req')
        all_params.append('_return_http_data_only')
        all_params.append('_preload_content')
        all_params.append('_request_timeout')

        params = locals()
        for key, val in six.iteritems(params['kwargs']):
            if key not in all_params:
                raise TypeError(
                    "Got an unexpected keyword argument '%s'"
                    " to method list" % key
                )
            params[key] = val
        del params['kwargs']
        # verify the required parameter 'list_request' is set
        if self.api_client.client_side_validation and ('list_request' not in params or
                                                       params['list_request'] is None):  # noqa: E501
            raise ValueError("Missing the required parameter `list_request` when calling `list`")  # noqa: E501

        collection_formats = {}

        path_params = {}

        query_params = []

        header_params = {}

        form_params = []
        local_var_files = {}

        body_params = None
        if 'list_request' in params:
            body_params = params['list_request']
        # HTTP header `Accept`
        header_params['Accept'] = self.api_client.select_header_accept(
            ['application/json'])  # noqa: E501

        # HTTP header `Content-Type`
        header_params['Content-Type'] = self.api_client.select_header_content_type(  # noqa: E501
            ['application/json'])  # noqa: E501

        # Authentication setting
        auth_settings = []  # noqa: E501

        return self.api_client.call_api(
            '/requester/list', 'POST',
            path_params,
            query_params,
            header_params,
            body=body_params,
            post_params=form_params,
            files=local_var_files,
            response_type='ListResponse',  # noqa: E501
            auth_settings=auth_settings,
            async_req=params.get('async_req'),
            _return_http_data_only=params.get('_return_http_data_only'),
            _preload_content=params.get('_preload_content', True),
            _request_timeout=params.get('_request_timeout'),
            collection_formats=collection_formats)

    def local_events(self, local_events_request, **kwargs):  # noqa: E501
        """Returns the node's local events related to the job-id passed in the body payload. Useful for troubleshooting.  # noqa: E501

        Local events (e.g. Selected, BidAccepted, Verified) are useful to track the progress of a job.  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True
        >>> thread = api.local_events(local_events_request, async_req=True)
        >>> result = thread.get()

        :param async_req bool
        :param LocalEventsRequest local_events_request:   (required)
        :return: LocalEventsResponse
                 If the method is called asynchronously,
                 returns the request thread.
        """
        kwargs['_return_http_data_only'] = True
        if kwargs.get('async_req'):
            return self.local_events_with_http_info(local_events_request, **kwargs)  # noqa: E501
        else:
            (data) = self.local_events_with_http_info(local_events_request, **kwargs)  # noqa: E501
            return data

    def local_events_with_http_info(self, local_events_request, **kwargs):  # noqa: E501
        """Returns the node's local events related to the job-id passed in the body payload. Useful for troubleshooting.  # noqa: E501

        Local events (e.g. Selected, BidAccepted, Verified) are useful to track the progress of a job.  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True
        >>> thread = api.local_events_with_http_info(local_events_request, async_req=True)
        >>> result = thread.get()

        :param async_req bool
        :param LocalEventsRequest local_events_request:   (required)
        :return: LocalEventsResponse
                 If the method is called asynchronously,
                 returns the request thread.
        """

        all_params = ['local_events_request']  # noqa: E501
        all_params.append('async_req')
        all_params.append('_return_http_data_only')
        all_params.append('_preload_content')
        all_params.append('_request_timeout')

        params = locals()
        for key, val in six.iteritems(params['kwargs']):
            if key not in all_params:
                raise TypeError(
                    "Got an unexpected keyword argument '%s'"
                    " to method local_events" % key
                )
            params[key] = val
        del params['kwargs']
        # verify the required parameter 'local_events_request' is set
        if self.api_client.client_side_validation and ('local_events_request' not in params or
                                                       params['local_events_request'] is None):  # noqa: E501
            raise ValueError("Missing the required parameter `local_events_request` when calling `local_events`")  # noqa: E501

        collection_formats = {}

        path_params = {}

        query_params = []

        header_params = {}

        form_params = []
        local_var_files = {}

        body_params = None
        if 'local_events_request' in params:
            body_params = params['local_events_request']
        # HTTP header `Accept`
        header_params['Accept'] = self.api_client.select_header_accept(
            ['application/json'])  # noqa: E501

        # HTTP header `Content-Type`
        header_params['Content-Type'] = self.api_client.select_header_content_type(  # noqa: E501
            ['application/json'])  # noqa: E501

        # Authentication setting
        auth_settings = []  # noqa: E501

        return self.api_client.call_api(
            '/requester/local_events', 'POST',
            path_params,
            query_params,
            header_params,
            body=body_params,
            post_params=form_params,
            files=local_var_files,
            response_type='LocalEventsResponse',  # noqa: E501
            auth_settings=auth_settings,
            async_req=params.get('async_req'),
            _return_http_data_only=params.get('_return_http_data_only'),
            _preload_content=params.get('_preload_content', True),
            _request_timeout=params.get('_request_timeout'),
            collection_formats=collection_formats)

    def results(self, state_request, **kwargs):  # noqa: E501
        """Returns the results of the job-id specified in the body payload.  # noqa: E501

        Example response:  ```json {   \"results\": [     {       \"NodeID\": \"QmdZQ7ZbhnvWY1J12XYKGHApJ6aufKyLNSvf8jZBrBaAVL\",       \"Data\": {         \"StorageSource\": \"IPFS\",         \"Name\": \"job-9304c616-291f-41ad-b862-54e133c0149e-shard-0-host-QmdZQ7ZbhnvWY1J12XYKGHApJ6aufKyLNSvf8jZBrBaAVL\",         \"CID\": \"QmTVmC7JBD2ES2qGPqBNVWnX1KeEPNrPGb7rJ8cpFgtefe\"       }     }   ] } ```  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True
        >>> thread = api.results(state_request, async_req=True)
        >>> result = thread.get()

        :param async_req bool
        :param StateRequest state_request:   (required)
        :return: ResultsResponse
                 If the method is called asynchronously,
                 returns the request thread.
        """
        kwargs['_return_http_data_only'] = True
        if kwargs.get('async_req'):
            return self.results_with_http_info(state_request, **kwargs)  # noqa: E501
        else:
            (data) = self.results_with_http_info(state_request, **kwargs)  # noqa: E501
            return data

    def results_with_http_info(self, state_request, **kwargs):  # noqa: E501
        """Returns the results of the job-id specified in the body payload.  # noqa: E501

        Example response:  ```json {   \"results\": [     {       \"NodeID\": \"QmdZQ7ZbhnvWY1J12XYKGHApJ6aufKyLNSvf8jZBrBaAVL\",       \"Data\": {         \"StorageSource\": \"IPFS\",         \"Name\": \"job-9304c616-291f-41ad-b862-54e133c0149e-shard-0-host-QmdZQ7ZbhnvWY1J12XYKGHApJ6aufKyLNSvf8jZBrBaAVL\",         \"CID\": \"QmTVmC7JBD2ES2qGPqBNVWnX1KeEPNrPGb7rJ8cpFgtefe\"       }     }   ] } ```  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True
        >>> thread = api.results_with_http_info(state_request, async_req=True)
        >>> result = thread.get()

        :param async_req bool
        :param StateRequest state_request:   (required)
        :return: ResultsResponse
                 If the method is called asynchronously,
                 returns the request thread.
        """

        all_params = ['state_request']  # noqa: E501
        all_params.append('async_req')
        all_params.append('_return_http_data_only')
        all_params.append('_preload_content')
        all_params.append('_request_timeout')

        params = locals()
        for key, val in six.iteritems(params['kwargs']):
            if key not in all_params:
                raise TypeError(
                    "Got an unexpected keyword argument '%s'"
                    " to method results" % key
                )
            params[key] = val
        del params['kwargs']
        # verify the required parameter 'state_request' is set
        if self.api_client.client_side_validation and ('state_request' not in params or
                                                       params['state_request'] is None):  # noqa: E501
            raise ValueError("Missing the required parameter `state_request` when calling `results`")  # noqa: E501

        collection_formats = {}

        path_params = {}

        query_params = []

        header_params = {}

        form_params = []
        local_var_files = {}

        body_params = None
        if 'state_request' in params:
            body_params = params['state_request']
        # HTTP header `Accept`
        header_params['Accept'] = self.api_client.select_header_accept(
            ['application/json'])  # noqa: E501

        # HTTP header `Content-Type`
        header_params['Content-Type'] = self.api_client.select_header_content_type(  # noqa: E501
            ['application/json'])  # noqa: E501

        # Authentication setting
        auth_settings = []  # noqa: E501

        return self.api_client.call_api(
            '/requester/results', 'POST',
            path_params,
            query_params,
            header_params,
            body=body_params,
            post_params=form_params,
            files=local_var_files,
            response_type='ResultsResponse',  # noqa: E501
            auth_settings=auth_settings,
            async_req=params.get('async_req'),
            _return_http_data_only=params.get('_return_http_data_only'),
            _preload_content=params.get('_preload_content', True),
            _request_timeout=params.get('_request_timeout'),
            collection_formats=collection_formats)

    def states(self, state_request, **kwargs):  # noqa: E501
        """Returns the state of the job-id specified in the body payload.  # noqa: E501

        Example response:  ```json {   \"state\": {     \"Nodes\": {       \"QmSyJ8VUd4YSPwZFJSJsHmmmmg7sd4BAc2yHY73nisJo86\": {         \"Shards\": {           \"0\": {             \"NodeId\": \"QmSyJ8VUd4YSPwZFJSJsHmmmmg7sd4BAc2yHY73nisJo86\",             \"State\": \"Cancelled\",             \"VerificationResult\": {},             \"PublishedResults\": {}           }         }       },       \"QmYgxZiySj3MRkwLSL4X2MF5F9f2PMhAE3LV49XkfNL1o3\": {         \"Shards\": {           \"0\": {             \"NodeId\": \"QmYgxZiySj3MRkwLSL4X2MF5F9f2PMhAE3LV49XkfNL1o3\",             \"State\": \"Cancelled\",             \"VerificationResult\": {},             \"PublishedResults\": {}           }         }       },       \"QmdZQ7ZbhnvWY1J12XYKGHApJ6aufKyLNSvf8jZBrBaAVL\": {         \"Shards\": {           \"0\": {             \"NodeId\": \"QmdZQ7ZbhnvWY1J12XYKGHApJ6aufKyLNSvf8jZBrBaAVL\",             \"State\": \"Completed\",             \"Status\": \"Got results proposal of length: 0\",             \"VerificationResult\": {               \"Complete\": true,               \"Result\": true             },             \"PublishedResults\": {               \"StorageSource\": \"IPFS\",               \"Name\": \"job-9304c616-291f-41ad-b862-54e133c0149e-shard-0-host-QmdZQ7ZbhnvWY1J12XYKGHApJ6aufKyLNSvf8jZBrBaAVL\",               \"CID\": \"QmTVmC7JBD2ES2qGPqBNVWnX1KeEPNrPGb7rJ8cpFgtefe\"             },             \"RunOutput\": {               \"stdout\": \"Thu Nov 17 13:32:55 UTC 2022\\n\",               \"stdouttruncated\": false,               \"stderr\": \"\",               \"stderrtruncated\": false,               \"exitCode\": 0,               \"runnerError\": \"\"             }           }         }       }     }   } } ```  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True
        >>> thread = api.states(state_request, async_req=True)
        >>> result = thread.get()

        :param async_req bool
        :param StateRequest state_request:   (required)
        :return: StateResponse
                 If the method is called asynchronously,
                 returns the request thread.
        """
        kwargs['_return_http_data_only'] = True
        if kwargs.get('async_req'):
            return self.states_with_http_info(state_request, **kwargs)  # noqa: E501
        else:
            (data) = self.states_with_http_info(state_request, **kwargs)  # noqa: E501
            return data

    def states_with_http_info(self, state_request, **kwargs):  # noqa: E501
        """Returns the state of the job-id specified in the body payload.  # noqa: E501

        Example response:  ```json {   \"state\": {     \"Nodes\": {       \"QmSyJ8VUd4YSPwZFJSJsHmmmmg7sd4BAc2yHY73nisJo86\": {         \"Shards\": {           \"0\": {             \"NodeId\": \"QmSyJ8VUd4YSPwZFJSJsHmmmmg7sd4BAc2yHY73nisJo86\",             \"State\": \"Cancelled\",             \"VerificationResult\": {},             \"PublishedResults\": {}           }         }       },       \"QmYgxZiySj3MRkwLSL4X2MF5F9f2PMhAE3LV49XkfNL1o3\": {         \"Shards\": {           \"0\": {             \"NodeId\": \"QmYgxZiySj3MRkwLSL4X2MF5F9f2PMhAE3LV49XkfNL1o3\",             \"State\": \"Cancelled\",             \"VerificationResult\": {},             \"PublishedResults\": {}           }         }       },       \"QmdZQ7ZbhnvWY1J12XYKGHApJ6aufKyLNSvf8jZBrBaAVL\": {         \"Shards\": {           \"0\": {             \"NodeId\": \"QmdZQ7ZbhnvWY1J12XYKGHApJ6aufKyLNSvf8jZBrBaAVL\",             \"State\": \"Completed\",             \"Status\": \"Got results proposal of length: 0\",             \"VerificationResult\": {               \"Complete\": true,               \"Result\": true             },             \"PublishedResults\": {               \"StorageSource\": \"IPFS\",               \"Name\": \"job-9304c616-291f-41ad-b862-54e133c0149e-shard-0-host-QmdZQ7ZbhnvWY1J12XYKGHApJ6aufKyLNSvf8jZBrBaAVL\",               \"CID\": \"QmTVmC7JBD2ES2qGPqBNVWnX1KeEPNrPGb7rJ8cpFgtefe\"             },             \"RunOutput\": {               \"stdout\": \"Thu Nov 17 13:32:55 UTC 2022\\n\",               \"stdouttruncated\": false,               \"stderr\": \"\",               \"stderrtruncated\": false,               \"exitCode\": 0,               \"runnerError\": \"\"             }           }         }       }     }   } } ```  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True
        >>> thread = api.states_with_http_info(state_request, async_req=True)
        >>> result = thread.get()

        :param async_req bool
        :param StateRequest state_request:   (required)
        :return: StateResponse
                 If the method is called asynchronously,
                 returns the request thread.
        """

        all_params = ['state_request']  # noqa: E501
        all_params.append('async_req')
        all_params.append('_return_http_data_only')
        all_params.append('_preload_content')
        all_params.append('_request_timeout')

        params = locals()
        for key, val in six.iteritems(params['kwargs']):
            if key not in all_params:
                raise TypeError(
                    "Got an unexpected keyword argument '%s'"
                    " to method states" % key
                )
            params[key] = val
        del params['kwargs']
        # verify the required parameter 'state_request' is set
        if self.api_client.client_side_validation and ('state_request' not in params or
                                                       params['state_request'] is None):  # noqa: E501
            raise ValueError("Missing the required parameter `state_request` when calling `states`")  # noqa: E501

        collection_formats = {}

        path_params = {}

        query_params = []

        header_params = {}

        form_params = []
        local_var_files = {}

        body_params = None
        if 'state_request' in params:
            body_params = params['state_request']
        # HTTP header `Accept`
        header_params['Accept'] = self.api_client.select_header_accept(
            ['application/json'])  # noqa: E501

        # HTTP header `Content-Type`
        header_params['Content-Type'] = self.api_client.select_header_content_type(  # noqa: E501
            ['application/json'])  # noqa: E501

        # Authentication setting
        auth_settings = []  # noqa: E501

        return self.api_client.call_api(
            '/requester/states', 'POST',
            path_params,
            query_params,
            header_params,
            body=body_params,
            post_params=form_params,
            files=local_var_files,
            response_type='StateResponse',  # noqa: E501
            auth_settings=auth_settings,
            async_req=params.get('async_req'),
            _return_http_data_only=params.get('_return_http_data_only'),
            _preload_content=params.get('_preload_content', True),
            _request_timeout=params.get('_request_timeout'),
            collection_formats=collection_formats)

    def submit(self, submit_request, **kwargs):  # noqa: E501
        """Submits a new job to the network.  # noqa: E501

        Description:  * `client_public_key`: The base64-encoded public key of the client. * `signature`: A base64-encoded signature of the `data` attribute, signed by the client. * `job_create_payload`:     * `ClientID`: Request must specify a `ClientID`. To retrieve your `ClientID`, you can do the following: (1) submit a dummy job to Bacalhau (or use one you created before), (2) run `bacalhau describe <job-id>` and fetch the `ClientID` field.  * `APIVersion`: e.g. `\"V1beta1\"`.     * `Spec`: https://github.com/bacalhau-project/bacalhau/blob/main/pkg/job.go   # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True
        >>> thread = api.submit(submit_request, async_req=True)
        >>> result = thread.get()

        :param async_req bool
        :param SubmitRequest submit_request:   (required)
        :return: SubmitResponse
                 If the method is called asynchronously,
                 returns the request thread.
        """
        kwargs['_return_http_data_only'] = True
        if kwargs.get('async_req'):
            return self.submit_with_http_info(submit_request, **kwargs)  # noqa: E501
        else:
            (data) = self.submit_with_http_info(submit_request, **kwargs)  # noqa: E501
            return data

    def submit_with_http_info(self, submit_request, **kwargs):  # noqa: E501
        """Submits a new job to the network.  # noqa: E501

        Description:  * `client_public_key`: The base64-encoded public key of the client. * `signature`: A base64-encoded signature of the `data` attribute, signed by the client. * `job_create_payload`:     * `ClientID`: Request must specify a `ClientID`. To retrieve your `ClientID`, you can do the following: (1) submit a dummy job to Bacalhau (or use one you created before), (2) run `bacalhau describe <job-id>` and fetch the `ClientID` field.  * `APIVersion`: e.g. `\"V1beta1\"`.     * `Spec`: https://github.com/bacalhau-project/bacalhau/blob/main/pkg/job.go   # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True
        >>> thread = api.submit_with_http_info(submit_request, async_req=True)
        >>> result = thread.get()

        :param async_req bool
        :param SubmitRequest submit_request:   (required)
        :return: SubmitResponse
                 If the method is called asynchronously,
                 returns the request thread.
        """

        all_params = ['submit_request']  # noqa: E501
        all_params.append('async_req')
        all_params.append('_return_http_data_only')
        all_params.append('_preload_content')
        all_params.append('_request_timeout')

        params = locals()
        for key, val in six.iteritems(params['kwargs']):
            if key not in all_params:
                raise TypeError(
                    "Got an unexpected keyword argument '%s'"
                    " to method submit" % key
                )
            params[key] = val
        del params['kwargs']
        # verify the required parameter 'submit_request' is set
        if self.api_client.client_side_validation and ('submit_request' not in params or
                                                       params['submit_request'] is None):  # noqa: E501
            raise ValueError("Missing the required parameter `submit_request` when calling `submit`")  # noqa: E501

        collection_formats = {}

        path_params = {}

        query_params = []

        header_params = {}

        form_params = []
        local_var_files = {}

        body_params = None
        if 'submit_request' in params:
            body_params = params['submit_request']
        # HTTP header `Accept`
        header_params['Accept'] = self.api_client.select_header_accept(
            ['application/json'])  # noqa: E501

        # HTTP header `Content-Type`
        header_params['Content-Type'] = self.api_client.select_header_content_type(  # noqa: E501
            ['application/json'])  # noqa: E501

        # Authentication setting
        auth_settings = []  # noqa: E501

        return self.api_client.call_api(
            '/requester/submit', 'POST',
            path_params,
            query_params,
            header_params,
            body=body_params,
            post_params=form_params,
            files=local_var_files,
            response_type='SubmitResponse',  # noqa: E501
            auth_settings=auth_settings,
            async_req=params.get('async_req'),
            _return_http_data_only=params.get('_return_http_data_only'),
            _preload_content=params.get('_preload_content', True),
            _request_timeout=params.get('_request_timeout'),
            collection_formats=collection_formats)
