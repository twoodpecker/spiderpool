// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package ip_test

import (
	"net"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/spidernet-io/spiderpool/pkg/constant"
	spiderpoolip "github.com/spidernet-io/spiderpool/pkg/ip"
)

var _ = Describe("CIDR", Label("cidr_test"), func() {
	Describe("Test ParseCIDR", func() {
		When("Verifying", func() {
			It("inputs invalid IP version", func() {
				ip, err := spiderpoolip.ParseCIDR(invalidIPVersion, "172.18.40.40/24")
				Expect(err).To(MatchError(spiderpoolip.ErrInvalidIPVersion))
				Expect(ip).To(BeNil())
			})

			It("inputs invalid CIDR address", func() {
				ip, err := spiderpoolip.ParseCIDR(constant.IPv4, invalidCIDR)
				Expect(err).To(MatchError(spiderpoolip.ErrInvalidCIDRFormat))
				Expect(ip).To(BeNil())
			})
		})

		It("parses IPv4 CIDR address", func() {
			ipNet, err := spiderpoolip.ParseCIDR(constant.IPv4, "172.18.40.40/24")
			Expect(err).NotTo(HaveOccurred())
			Expect(ipNet).To(Equal(
				&net.IPNet{
					IP:   net.IPv4(172, 18, 40, 0).To4(),
					Mask: net.CIDRMask(24, 32),
				},
			))
		})

		It("parses IPv6 CIDR address", func() {
			ipNet, err := spiderpoolip.ParseCIDR(constant.IPv6, "abcd:1234::1/120")
			Expect(err).NotTo(HaveOccurred())
			Expect(ipNet).To(Equal(
				&net.IPNet{
					IP:   net.ParseIP("abcd:1234::0"),
					Mask: net.CIDRMask(120, 128),
				},
			))
		})
	})

	Describe("Test ContainsCIDR", func() {
		When("Verifying", func() {
			It("inputs invalid IP version", func() {
				contains, err := spiderpoolip.ContainsCIDR(invalidIPVersion, "172.18.40.0/24", "172.18.40.0/25")
				Expect(err).To(MatchError(spiderpoolip.ErrInvalidIPVersion))
				Expect(contains).To(BeFalse())
			})

			It("inputs invalid CIDR address", func() {
				contains, err := spiderpoolip.ContainsCIDR(constant.IPv4, invalidCIDR, "172.18.40.0/25")
				Expect(err).To(MatchError(spiderpoolip.ErrInvalidCIDRFormat))
				Expect(contains).To(BeFalse())

				contains, err = spiderpoolip.ContainsCIDR(constant.IPv4, "172.18.40.0/24", invalidCIDR)
				Expect(err).To(MatchError(spiderpoolip.ErrInvalidCIDRFormat))
				Expect(contains).To(BeFalse())
			})
		})

		When("IPv4", func() {
			It("tests that a CIDR address contains another CIDR address", func() {
				contains, err := spiderpoolip.ContainsCIDR(constant.IPv4, "172.18.40.0/24", "172.18.40.0/25")
				Expect(err).NotTo(HaveOccurred())
				Expect(contains).To(BeTrue())
			})

			It("tests that one CIDR address does not contain another CIDR address", func() {
				contains, err := spiderpoolip.ContainsCIDR(constant.IPv4, "172.18.40.0/25", "172.18.40.0/24")
				Expect(err).NotTo(HaveOccurred())
				Expect(contains).To(BeFalse())
			})
		})

		When("IPv6", func() {
			It("tests that a CIDR address contains another CIDR address", func() {
				contains, err := spiderpoolip.ContainsCIDR(constant.IPv6, "abcd:1234::/120", "abcd:1234::/121")
				Expect(err).NotTo(HaveOccurred())
				Expect(contains).To(BeTrue())
			})

			It("tests that one CIDR address does not contain another CIDR address", func() {
				contains, err := spiderpoolip.ContainsCIDR(constant.IPv6, "abcd:1234::/121", "abcd:1234::/120")
				Expect(err).NotTo(HaveOccurred())
				Expect(contains).To(BeFalse())
			})
		})
	})

	Describe("Test IsCIDROverlap", func() {
		When("Verifying", func() {
			It("inputs invalid IP version", func() {
				overlap, err := spiderpoolip.IsCIDROverlap(invalidIPVersion, "172.18.40.0/24", "172.18.40.0/25")
				Expect(err).To(MatchError(spiderpoolip.ErrInvalidIPVersion))
				Expect(overlap).To(BeFalse())
			})

			It("inputs invalid CIDR address", func() {
				overlap, err := spiderpoolip.IsCIDROverlap(constant.IPv4, invalidCIDR, "172.18.40.0/25")
				Expect(err).To(MatchError(spiderpoolip.ErrInvalidCIDRFormat))
				Expect(overlap).To(BeFalse())

				overlap, err = spiderpoolip.IsCIDROverlap(constant.IPv4, "172.18.40.0/24", invalidCIDR)
				Expect(err).To(MatchError(spiderpoolip.ErrInvalidCIDRFormat))
				Expect(overlap).To(BeFalse())
			})
		})

		When("IPv4", func() {
			It("tests that two CIDR addresses overlap", func() {
				overlap, err := spiderpoolip.IsCIDROverlap(constant.IPv4, "172.18.40.0/24", "172.18.40.0/25")
				Expect(err).NotTo(HaveOccurred())
				Expect(overlap).To(BeTrue())
			})

			It("tests that two CIDR addresses do not overlap", func() {
				overlap, err := spiderpoolip.IsCIDROverlap(constant.IPv4, "172.18.41.0/24", "172.18.40.0/24")
				Expect(err).NotTo(HaveOccurred())
				Expect(overlap).To(BeFalse())
			})
		})

		When("IPv6", func() {
			It("tests that two CIDR addresses overlap", func() {
				overlap, err := spiderpoolip.IsCIDROverlap(constant.IPv6, "abcd:1234::/120", "abcd:1234::/121")
				Expect(err).NotTo(HaveOccurred())
				Expect(overlap).To(BeTrue())
			})

			It("tests that two CIDR addresses do not overlap", func() {
				overlap, err := spiderpoolip.IsCIDROverlap(constant.IPv6, "abcd:1235::/120", "abcd:1234::/120")
				Expect(err).NotTo(HaveOccurred())
				Expect(overlap).To(BeFalse())
			})
		})
	})

	Describe("Test IsCIDR", func() {
		When("Verifying", func() {
			It("inputs invalid IP version", func() {
				err := spiderpoolip.IsCIDR(invalidIPVersion, "172.18.40.0/24")
				Expect(err).To(MatchError(spiderpoolip.ErrInvalidIPVersion))
			})

			It("inputs invalid CIDR address", func() {
				err := spiderpoolip.IsCIDR(constant.IPv4, invalidCIDR)
				Expect(err).To(MatchError(spiderpoolip.ErrInvalidCIDRFormat))
			})
		})

		It("is an IPv4 CIDR address", func() {
			Expect(spiderpoolip.IsCIDR(constant.IPv4, "172.18.40.0/24")).To(Succeed())
		})

		It("is an IPv6 CIDR address", func() {
			Expect(spiderpoolip.IsCIDR(constant.IPv6, "abcd:1234::/120")).To(Succeed())
		})
	})

	Describe("Test IsIPv4CIDR", func() {
		It("tests whether it is an IPv4 CIDR address", func() {
			Expect(spiderpoolip.IsIPv4CIDR(invalidCIDR)).To(BeFalse())
			Expect(spiderpoolip.IsIPv4CIDR("172.18.40.0/24")).To(BeTrue())
		})
	})

	Describe("Test IsIPv6CIDR", func() {
		It("tests whether it is an IPv6 CIDR address", func() {
			Expect(spiderpoolip.IsIPv6CIDR(invalidCIDR)).To(BeFalse())
			Expect(spiderpoolip.IsIPv6CIDR("abcd:1234::/120")).To(BeTrue())
		})
	})
})
