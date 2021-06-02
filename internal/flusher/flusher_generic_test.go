package flusher_test

import (
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/ozoncp/ocp-course-api/internal/flusher"
	"github.com/ozoncp/ocp-course-api/internal/mock_repo"
	"github.com/ozoncp/ocp-course-api/internal/repo"
	"github.com/ozoncp/ocp-course-api/internal/utils/commons"
)

var _ = Describe("FlusherGeneric", func() {
	const (
		batchSize int = 2
	)

	var (
		ctrl     *gomock.Controller
		mockRepo *mock_repo.MockRepoTValue
		flusher  FlusherTValue
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mock_repo.NewMockRepoTValue(ctrl)
	})

	JustBeforeEach(func() {
		flusher = NewFlusherTValue(mockRepo, commons.NewNaturalIntPanic(batchSize))
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("Flush data in a repo", func() {

		Context("all saved successfully", func() {

			It("call add to repo once", func() {
				data := []repo.TValue{1, 2}
				mockRepo.EXPECT().AddTValues([]TValue{1, 2}).Return(nil)
				rest := flusher.FlushTValue(data)
				Expect(rest).Should(BeNil())
			})

			It("call add to repo twice", func() {
				data := []repo.TValue{1, 2, 3}
				gomock.InOrder(
					mockRepo.EXPECT().AddTValues([]TValue{1, 2}).Return(nil),
					mockRepo.EXPECT().AddTValues([]TValue{3}).Return(nil),
				)
				rest := flusher.FlushTValue(data)
				Expect(rest).Should(BeNil())
			})
		})

		Context("some saving failed", func() {

			It("saving fails immediately", func() {
				data := []repo.TValue{1, 2, 3}
				mockRepo.EXPECT().AddTValues(gomock.Any()).Return(errors.New("BOOM!"))
				rest := flusher.FlushTValue(data)
				Expect(rest).Should(Equal(data[0:]))
			})

			It("the second saving fails", func() {
				data := []repo.TValue{1, 2, 3}
				gomock.InOrder(
					mockRepo.EXPECT().AddTValues(gomock.Any()).Return(nil),
					mockRepo.EXPECT().AddTValues(gomock.Any()).Return(errors.New("BOOM!")),
				)
				rest := flusher.FlushTValue(data)
				Expect(rest).Should(Equal(data[2:]))
			})
		})
	})
})
