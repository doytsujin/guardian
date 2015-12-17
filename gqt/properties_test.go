package gqt_test

import (
	"github.com/cloudfoundry-incubator/garden"
	"github.com/cloudfoundry-incubator/guardian/gqt/runner"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Properties", func() {
	var (
		client    *runner.RunningGarden
		container garden.Container
		props     garden.Properties
	)

	BeforeEach(func() {
		var err error
		client = startGarden()
		props = garden.Properties{"somename": "somevalue"}
		container, err = client.Create(garden.ContainerSpec{
			Properties: props,
		})
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		client.DestroyAndStop()
	})

	It("can get properties", func() {
		properties, err := container.Properties()
		Expect(err).NotTo(HaveOccurred())
		Expect(properties).To(HaveKeyWithValue("somename", "somevalue"))
	})

	It("can set a single property", func() {
		err := container.SetProperty("someothername", "someothervalue")
		Expect(err).NotTo(HaveOccurred())

		properties, err := container.Properties()
		Expect(err).NotTo(HaveOccurred())
		Expect(properties).To(HaveKeyWithValue("somename", "somevalue"))
		Expect(properties).To(HaveKeyWithValue("someothername", "someothervalue"))
	})

	It("can get a single property", func() {
		err := container.SetProperty("bing", "bong")
		Expect(err).NotTo(HaveOccurred())

		value, err := container.Property("bing")
		Expect(err).NotTo(HaveOccurred())
		Expect(value).To(Equal("bong"))
	})

	It("can remove a single property", func() {
		err := container.SetProperty("bing", "bong")
		Expect(err).NotTo(HaveOccurred())

		err = container.RemoveProperty("bing")
		Expect(err).NotTo(HaveOccurred())

		_, err = container.Property("bing")
		Expect(err).To(HaveOccurred())
	})

	It("can filter containers based on their properties", func() {
		_, err := client.Create(garden.ContainerSpec{
			Properties: garden.Properties{
				"somename": "wrongvalue",
			},
		})
		Expect(err).NotTo(HaveOccurred())

		containers, err := client.Containers(props)
		Expect(err).NotTo(HaveOccurred())
		Expect(containers).To(HaveLen(1))
		Expect(containers).To(ConsistOf(container))
	})
})