package internalversion

import (
	api "github.com/openshift/origin/pkg/image/api"
	scheme "github.com/openshift/origin/pkg/image/generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ImageStreamsGetter has a method to return a ImageStreamInterface.
// A group's client should implement this interface.
type ImageStreamsGetter interface {
	ImageStreams(namespace string) ImageStreamInterface
}

// ImageStreamInterface has methods to work with ImageStream resources.
type ImageStreamInterface interface {
	Create(*api.ImageStream) (*api.ImageStream, error)
	Update(*api.ImageStream) (*api.ImageStream, error)
	UpdateStatus(*api.ImageStream) (*api.ImageStream, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*api.ImageStream, error)
	List(opts v1.ListOptions) (*api.ImageStreamList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *api.ImageStream, err error)
	ImageStreamExpansion
}

// imageStreams implements ImageStreamInterface
type imageStreams struct {
	client rest.Interface
	ns     string
}

// newImageStreams returns a ImageStreams
func newImageStreams(c *ImageClient, namespace string) *imageStreams {
	return &imageStreams{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a imageStream and creates it.  Returns the server's representation of the imageStream, and an error, if there is any.
func (c *imageStreams) Create(imageStream *api.ImageStream) (result *api.ImageStream, err error) {
	result = &api.ImageStream{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("imagestreams").
		Body(imageStream).
		Do().
		Into(result)
	return
}

// Update takes the representation of a imageStream and updates it. Returns the server's representation of the imageStream, and an error, if there is any.
func (c *imageStreams) Update(imageStream *api.ImageStream) (result *api.ImageStream, err error) {
	result = &api.ImageStream{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("imagestreams").
		Name(imageStream.Name).
		Body(imageStream).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclientstatus=false comment above the type to avoid generating UpdateStatus().

func (c *imageStreams) UpdateStatus(imageStream *api.ImageStream) (result *api.ImageStream, err error) {
	result = &api.ImageStream{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("imagestreams").
		Name(imageStream.Name).
		SubResource("status").
		Body(imageStream).
		Do().
		Into(result)
	return
}

// Delete takes name of the imageStream and deletes it. Returns an error if one occurs.
func (c *imageStreams) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("imagestreams").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *imageStreams) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("imagestreams").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the imageStream, and returns the corresponding imageStream object, and an error if there is any.
func (c *imageStreams) Get(name string, options v1.GetOptions) (result *api.ImageStream, err error) {
	result = &api.ImageStream{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("imagestreams").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ImageStreams that match those selectors.
func (c *imageStreams) List(opts v1.ListOptions) (result *api.ImageStreamList, err error) {
	result = &api.ImageStreamList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("imagestreams").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested imageStreams.
func (c *imageStreams) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("imagestreams").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched imageStream.
func (c *imageStreams) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *api.ImageStream, err error) {
	result = &api.ImageStream{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("imagestreams").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
