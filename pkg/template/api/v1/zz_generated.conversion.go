// +build !ignore_autogenerated_openshift

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1

import (
	api "github.com/openshift/origin/pkg/template/api"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	pkg_api "k8s.io/kubernetes/pkg/api"
	api_v1 "k8s.io/kubernetes/pkg/api/v1"
	unsafe "unsafe"
)

func init() {
	SchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1_BrokerTemplateInstance_To_api_BrokerTemplateInstance,
		Convert_api_BrokerTemplateInstance_To_v1_BrokerTemplateInstance,
		Convert_v1_BrokerTemplateInstanceList_To_api_BrokerTemplateInstanceList,
		Convert_api_BrokerTemplateInstanceList_To_v1_BrokerTemplateInstanceList,
		Convert_v1_BrokerTemplateInstanceSpec_To_api_BrokerTemplateInstanceSpec,
		Convert_api_BrokerTemplateInstanceSpec_To_v1_BrokerTemplateInstanceSpec,
		Convert_v1_Parameter_To_api_Parameter,
		Convert_api_Parameter_To_v1_Parameter,
		Convert_v1_Template_To_api_Template,
		Convert_api_Template_To_v1_Template,
		Convert_v1_TemplateInstance_To_api_TemplateInstance,
		Convert_api_TemplateInstance_To_v1_TemplateInstance,
		Convert_v1_TemplateInstanceCondition_To_api_TemplateInstanceCondition,
		Convert_api_TemplateInstanceCondition_To_v1_TemplateInstanceCondition,
		Convert_v1_TemplateInstanceList_To_api_TemplateInstanceList,
		Convert_api_TemplateInstanceList_To_v1_TemplateInstanceList,
		Convert_v1_TemplateInstanceRequester_To_api_TemplateInstanceRequester,
		Convert_api_TemplateInstanceRequester_To_v1_TemplateInstanceRequester,
		Convert_v1_TemplateInstanceSpec_To_api_TemplateInstanceSpec,
		Convert_api_TemplateInstanceSpec_To_v1_TemplateInstanceSpec,
		Convert_v1_TemplateInstanceStatus_To_api_TemplateInstanceStatus,
		Convert_api_TemplateInstanceStatus_To_v1_TemplateInstanceStatus,
		Convert_v1_TemplateList_To_api_TemplateList,
		Convert_api_TemplateList_To_v1_TemplateList,
	)
}

func autoConvert_v1_BrokerTemplateInstance_To_api_BrokerTemplateInstance(in *BrokerTemplateInstance, out *api.BrokerTemplateInstance, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_BrokerTemplateInstanceSpec_To_api_BrokerTemplateInstanceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_BrokerTemplateInstance_To_api_BrokerTemplateInstance(in *BrokerTemplateInstance, out *api.BrokerTemplateInstance, s conversion.Scope) error {
	return autoConvert_v1_BrokerTemplateInstance_To_api_BrokerTemplateInstance(in, out, s)
}

func autoConvert_api_BrokerTemplateInstance_To_v1_BrokerTemplateInstance(in *api.BrokerTemplateInstance, out *BrokerTemplateInstance, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_api_BrokerTemplateInstanceSpec_To_v1_BrokerTemplateInstanceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

func Convert_api_BrokerTemplateInstance_To_v1_BrokerTemplateInstance(in *api.BrokerTemplateInstance, out *BrokerTemplateInstance, s conversion.Scope) error {
	return autoConvert_api_BrokerTemplateInstance_To_v1_BrokerTemplateInstance(in, out, s)
}

func autoConvert_v1_BrokerTemplateInstanceList_To_api_BrokerTemplateInstanceList(in *BrokerTemplateInstanceList, out *api.BrokerTemplateInstanceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]api.BrokerTemplateInstance, len(*in))
		for i := range *in {
			if err := Convert_v1_BrokerTemplateInstance_To_api_BrokerTemplateInstance(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v1_BrokerTemplateInstanceList_To_api_BrokerTemplateInstanceList(in *BrokerTemplateInstanceList, out *api.BrokerTemplateInstanceList, s conversion.Scope) error {
	return autoConvert_v1_BrokerTemplateInstanceList_To_api_BrokerTemplateInstanceList(in, out, s)
}

func autoConvert_api_BrokerTemplateInstanceList_To_v1_BrokerTemplateInstanceList(in *api.BrokerTemplateInstanceList, out *BrokerTemplateInstanceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BrokerTemplateInstance, len(*in))
		for i := range *in {
			if err := Convert_api_BrokerTemplateInstance_To_v1_BrokerTemplateInstance(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = make([]BrokerTemplateInstance, 0)
	}
	return nil
}

func Convert_api_BrokerTemplateInstanceList_To_v1_BrokerTemplateInstanceList(in *api.BrokerTemplateInstanceList, out *BrokerTemplateInstanceList, s conversion.Scope) error {
	return autoConvert_api_BrokerTemplateInstanceList_To_v1_BrokerTemplateInstanceList(in, out, s)
}

func autoConvert_v1_BrokerTemplateInstanceSpec_To_api_BrokerTemplateInstanceSpec(in *BrokerTemplateInstanceSpec, out *api.BrokerTemplateInstanceSpec, s conversion.Scope) error {
	if err := api_v1.Convert_v1_ObjectReference_To_api_ObjectReference(&in.TemplateInstance, &out.TemplateInstance, s); err != nil {
		return err
	}
	if err := api_v1.Convert_v1_ObjectReference_To_api_ObjectReference(&in.Secret, &out.Secret, s); err != nil {
		return err
	}
	out.BindingIDs = *(*[]string)(unsafe.Pointer(&in.BindingIDs))
	return nil
}

func Convert_v1_BrokerTemplateInstanceSpec_To_api_BrokerTemplateInstanceSpec(in *BrokerTemplateInstanceSpec, out *api.BrokerTemplateInstanceSpec, s conversion.Scope) error {
	return autoConvert_v1_BrokerTemplateInstanceSpec_To_api_BrokerTemplateInstanceSpec(in, out, s)
}

func autoConvert_api_BrokerTemplateInstanceSpec_To_v1_BrokerTemplateInstanceSpec(in *api.BrokerTemplateInstanceSpec, out *BrokerTemplateInstanceSpec, s conversion.Scope) error {
	if err := api_v1.Convert_api_ObjectReference_To_v1_ObjectReference(&in.TemplateInstance, &out.TemplateInstance, s); err != nil {
		return err
	}
	if err := api_v1.Convert_api_ObjectReference_To_v1_ObjectReference(&in.Secret, &out.Secret, s); err != nil {
		return err
	}
	if in.BindingIDs == nil {
		out.BindingIDs = make([]string, 0)
	} else {
		out.BindingIDs = *(*[]string)(unsafe.Pointer(&in.BindingIDs))
	}
	return nil
}

func Convert_api_BrokerTemplateInstanceSpec_To_v1_BrokerTemplateInstanceSpec(in *api.BrokerTemplateInstanceSpec, out *BrokerTemplateInstanceSpec, s conversion.Scope) error {
	return autoConvert_api_BrokerTemplateInstanceSpec_To_v1_BrokerTemplateInstanceSpec(in, out, s)
}

func autoConvert_v1_Parameter_To_api_Parameter(in *Parameter, out *api.Parameter, s conversion.Scope) error {
	out.Name = in.Name
	out.DisplayName = in.DisplayName
	out.Description = in.Description
	out.Value = in.Value
	out.Generate = in.Generate
	out.From = in.From
	out.Required = in.Required
	return nil
}

func Convert_v1_Parameter_To_api_Parameter(in *Parameter, out *api.Parameter, s conversion.Scope) error {
	return autoConvert_v1_Parameter_To_api_Parameter(in, out, s)
}

func autoConvert_api_Parameter_To_v1_Parameter(in *api.Parameter, out *Parameter, s conversion.Scope) error {
	out.Name = in.Name
	out.DisplayName = in.DisplayName
	out.Description = in.Description
	out.Value = in.Value
	out.Generate = in.Generate
	out.From = in.From
	out.Required = in.Required
	return nil
}

func Convert_api_Parameter_To_v1_Parameter(in *api.Parameter, out *Parameter, s conversion.Scope) error {
	return autoConvert_api_Parameter_To_v1_Parameter(in, out, s)
}

func autoConvert_v1_Template_To_api_Template(in *Template, out *api.Template, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Message = in.Message
	if in.Objects != nil {
		in, out := &in.Objects, &out.Objects
		*out = make([]runtime.Object, len(*in))
		for i := range *in {
			if err := runtime.Convert_runtime_RawExtension_To_runtime_Object(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Objects = nil
	}
	out.Parameters = *(*[]api.Parameter)(unsafe.Pointer(&in.Parameters))
	out.ObjectLabels = *(*map[string]string)(unsafe.Pointer(&in.ObjectLabels))
	return nil
}

func Convert_v1_Template_To_api_Template(in *Template, out *api.Template, s conversion.Scope) error {
	return autoConvert_v1_Template_To_api_Template(in, out, s)
}

func autoConvert_api_Template_To_v1_Template(in *api.Template, out *Template, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Message = in.Message
	out.Parameters = *(*[]Parameter)(unsafe.Pointer(&in.Parameters))
	if in.Objects != nil {
		in, out := &in.Objects, &out.Objects
		*out = make([]runtime.RawExtension, len(*in))
		for i := range *in {
			if err := runtime.Convert_runtime_Object_To_runtime_RawExtension(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Objects = make([]runtime.RawExtension, 0)
	}
	out.ObjectLabels = *(*map[string]string)(unsafe.Pointer(&in.ObjectLabels))
	return nil
}

func Convert_api_Template_To_v1_Template(in *api.Template, out *Template, s conversion.Scope) error {
	return autoConvert_api_Template_To_v1_Template(in, out, s)
}

func autoConvert_v1_TemplateInstance_To_api_TemplateInstance(in *TemplateInstance, out *api.TemplateInstance, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_TemplateInstanceSpec_To_api_TemplateInstanceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_TemplateInstanceStatus_To_api_TemplateInstanceStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_TemplateInstance_To_api_TemplateInstance(in *TemplateInstance, out *api.TemplateInstance, s conversion.Scope) error {
	return autoConvert_v1_TemplateInstance_To_api_TemplateInstance(in, out, s)
}

func autoConvert_api_TemplateInstance_To_v1_TemplateInstance(in *api.TemplateInstance, out *TemplateInstance, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_api_TemplateInstanceSpec_To_v1_TemplateInstanceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_api_TemplateInstanceStatus_To_v1_TemplateInstanceStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_api_TemplateInstance_To_v1_TemplateInstance(in *api.TemplateInstance, out *TemplateInstance, s conversion.Scope) error {
	return autoConvert_api_TemplateInstance_To_v1_TemplateInstance(in, out, s)
}

func autoConvert_v1_TemplateInstanceCondition_To_api_TemplateInstanceCondition(in *TemplateInstanceCondition, out *api.TemplateInstanceCondition, s conversion.Scope) error {
	out.Type = api.TemplateInstanceConditionType(in.Type)
	out.Status = pkg_api.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

func Convert_v1_TemplateInstanceCondition_To_api_TemplateInstanceCondition(in *TemplateInstanceCondition, out *api.TemplateInstanceCondition, s conversion.Scope) error {
	return autoConvert_v1_TemplateInstanceCondition_To_api_TemplateInstanceCondition(in, out, s)
}

func autoConvert_api_TemplateInstanceCondition_To_v1_TemplateInstanceCondition(in *api.TemplateInstanceCondition, out *TemplateInstanceCondition, s conversion.Scope) error {
	out.Type = TemplateInstanceConditionType(in.Type)
	out.Status = api_v1.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

func Convert_api_TemplateInstanceCondition_To_v1_TemplateInstanceCondition(in *api.TemplateInstanceCondition, out *TemplateInstanceCondition, s conversion.Scope) error {
	return autoConvert_api_TemplateInstanceCondition_To_v1_TemplateInstanceCondition(in, out, s)
}

func autoConvert_v1_TemplateInstanceList_To_api_TemplateInstanceList(in *TemplateInstanceList, out *api.TemplateInstanceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]api.TemplateInstance, len(*in))
		for i := range *in {
			if err := Convert_v1_TemplateInstance_To_api_TemplateInstance(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v1_TemplateInstanceList_To_api_TemplateInstanceList(in *TemplateInstanceList, out *api.TemplateInstanceList, s conversion.Scope) error {
	return autoConvert_v1_TemplateInstanceList_To_api_TemplateInstanceList(in, out, s)
}

func autoConvert_api_TemplateInstanceList_To_v1_TemplateInstanceList(in *api.TemplateInstanceList, out *TemplateInstanceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TemplateInstance, len(*in))
		for i := range *in {
			if err := Convert_api_TemplateInstance_To_v1_TemplateInstance(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = make([]TemplateInstance, 0)
	}
	return nil
}

func Convert_api_TemplateInstanceList_To_v1_TemplateInstanceList(in *api.TemplateInstanceList, out *TemplateInstanceList, s conversion.Scope) error {
	return autoConvert_api_TemplateInstanceList_To_v1_TemplateInstanceList(in, out, s)
}

func autoConvert_v1_TemplateInstanceRequester_To_api_TemplateInstanceRequester(in *TemplateInstanceRequester, out *api.TemplateInstanceRequester, s conversion.Scope) error {
	out.Username = in.Username
	return nil
}

func Convert_v1_TemplateInstanceRequester_To_api_TemplateInstanceRequester(in *TemplateInstanceRequester, out *api.TemplateInstanceRequester, s conversion.Scope) error {
	return autoConvert_v1_TemplateInstanceRequester_To_api_TemplateInstanceRequester(in, out, s)
}

func autoConvert_api_TemplateInstanceRequester_To_v1_TemplateInstanceRequester(in *api.TemplateInstanceRequester, out *TemplateInstanceRequester, s conversion.Scope) error {
	out.Username = in.Username
	return nil
}

func Convert_api_TemplateInstanceRequester_To_v1_TemplateInstanceRequester(in *api.TemplateInstanceRequester, out *TemplateInstanceRequester, s conversion.Scope) error {
	return autoConvert_api_TemplateInstanceRequester_To_v1_TemplateInstanceRequester(in, out, s)
}

func autoConvert_v1_TemplateInstanceSpec_To_api_TemplateInstanceSpec(in *TemplateInstanceSpec, out *api.TemplateInstanceSpec, s conversion.Scope) error {
	if err := Convert_v1_Template_To_api_Template(&in.Template, &out.Template, s); err != nil {
		return err
	}
	if err := api_v1.Convert_v1_LocalObjectReference_To_api_LocalObjectReference(&in.Secret, &out.Secret, s); err != nil {
		return err
	}
	out.Requester = (*api.TemplateInstanceRequester)(unsafe.Pointer(in.Requester))
	return nil
}

func Convert_v1_TemplateInstanceSpec_To_api_TemplateInstanceSpec(in *TemplateInstanceSpec, out *api.TemplateInstanceSpec, s conversion.Scope) error {
	return autoConvert_v1_TemplateInstanceSpec_To_api_TemplateInstanceSpec(in, out, s)
}

func autoConvert_api_TemplateInstanceSpec_To_v1_TemplateInstanceSpec(in *api.TemplateInstanceSpec, out *TemplateInstanceSpec, s conversion.Scope) error {
	if err := Convert_api_Template_To_v1_Template(&in.Template, &out.Template, s); err != nil {
		return err
	}
	if err := api_v1.Convert_api_LocalObjectReference_To_v1_LocalObjectReference(&in.Secret, &out.Secret, s); err != nil {
		return err
	}
	out.Requester = (*TemplateInstanceRequester)(unsafe.Pointer(in.Requester))
	return nil
}

func Convert_api_TemplateInstanceSpec_To_v1_TemplateInstanceSpec(in *api.TemplateInstanceSpec, out *TemplateInstanceSpec, s conversion.Scope) error {
	return autoConvert_api_TemplateInstanceSpec_To_v1_TemplateInstanceSpec(in, out, s)
}

func autoConvert_v1_TemplateInstanceStatus_To_api_TemplateInstanceStatus(in *TemplateInstanceStatus, out *api.TemplateInstanceStatus, s conversion.Scope) error {
	out.Conditions = *(*[]api.TemplateInstanceCondition)(unsafe.Pointer(&in.Conditions))
	return nil
}

func Convert_v1_TemplateInstanceStatus_To_api_TemplateInstanceStatus(in *TemplateInstanceStatus, out *api.TemplateInstanceStatus, s conversion.Scope) error {
	return autoConvert_v1_TemplateInstanceStatus_To_api_TemplateInstanceStatus(in, out, s)
}

func autoConvert_api_TemplateInstanceStatus_To_v1_TemplateInstanceStatus(in *api.TemplateInstanceStatus, out *TemplateInstanceStatus, s conversion.Scope) error {
	if in.Conditions == nil {
		out.Conditions = make([]TemplateInstanceCondition, 0)
	} else {
		out.Conditions = *(*[]TemplateInstanceCondition)(unsafe.Pointer(&in.Conditions))
	}
	return nil
}

func Convert_api_TemplateInstanceStatus_To_v1_TemplateInstanceStatus(in *api.TemplateInstanceStatus, out *TemplateInstanceStatus, s conversion.Scope) error {
	return autoConvert_api_TemplateInstanceStatus_To_v1_TemplateInstanceStatus(in, out, s)
}

func autoConvert_v1_TemplateList_To_api_TemplateList(in *TemplateList, out *api.TemplateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]api.Template, len(*in))
		for i := range *in {
			if err := Convert_v1_Template_To_api_Template(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v1_TemplateList_To_api_TemplateList(in *TemplateList, out *api.TemplateList, s conversion.Scope) error {
	return autoConvert_v1_TemplateList_To_api_TemplateList(in, out, s)
}

func autoConvert_api_TemplateList_To_v1_TemplateList(in *api.TemplateList, out *TemplateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Template, len(*in))
		for i := range *in {
			if err := Convert_api_Template_To_v1_Template(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = make([]Template, 0)
	}
	return nil
}

func Convert_api_TemplateList_To_v1_TemplateList(in *api.TemplateList, out *TemplateList, s conversion.Scope) error {
	return autoConvert_api_TemplateList_To_v1_TemplateList(in, out, s)
}
