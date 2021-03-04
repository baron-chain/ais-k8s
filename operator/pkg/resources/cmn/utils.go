// Package cmn provides utilities for common AIS cluster resources
/*
 * Copyright (c) 2021, NVIDIA CORPORATION. All rights reserved.
 */
package cmn

import (
	corev1 "k8s.io/api/core/v1"
)

func EnvFromFieldPath(envName, path string) corev1.EnvVar {
	return corev1.EnvVar{
		Name: envName,
		ValueFrom: &corev1.EnvVarSource{
			FieldRef: &corev1.ObjectFieldSelector{
				FieldPath: path,
			},
		},
	}
}

func EnvFromValue(envName, value string) corev1.EnvVar {
	return corev1.EnvVar{
		Name:  envName,
		Value: value,
	}
}

// IsBoolSet checks if a boolean pointer is set to true.
func IsBoolSet(v *bool) bool {
	return v != nil && *v
}
