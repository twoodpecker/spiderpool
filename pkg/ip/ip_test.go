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

var _ = Describe("IP", Label("ip_test"), func() {
	Describe("Test IsIPVersion", func() {
		It("inputs invalid IP version", func() {
			err := spiderpoolip.IsIPVersion(invalidIPVersion)
			Expect(err).To(MatchError(spiderpoolip.ErrInvalidIPVersion))
		})

		It("is IPv4", func() {
			Expect(spiderpoolip.IsIPVersion(constant.IPv4)).To(Succeed())
		})

		It("is IPv6", func() {
			Expect(spiderpoolip.IsIPVersion(constant.IPv6)).To(Succeed())
		})
	})

	Describe("Test ParseIP", func() {
		Describe("IP format", func() {
			When("Verifying", func() {
				It("inputs invalid IP version", func() {
					ip, err := spiderpoolip.ParseIP(invalidIPVersion, "172.18.40.40", false)
					Expect(err).To(MatchError(spiderpoolip.ErrInvalidIPVersion))
					Expect(ip).To(BeNil())
				})

				It("inputs invalid IP address", func() {
					ip, err := spiderpoolip.ParseIP(constant.IPv4, invalidIP, false)
					Expect(err).To(MatchError(spiderpoolip.ErrInvalidIPFormat))
					Expect(ip).To(BeNil())
				})
			})

			It("parses IPv4 IP address", func() {
				ip, err := spiderpoolip.ParseIP(constant.IPv4, "172.18.40.40", false)
				Expect(err).NotTo(HaveOccurred())
				Expect(ip).To(Equal(
					&net.IPNet{
						IP:   net.IPv4(172, 18, 40, 40),
						Mask: net.CIDRMask(32, 32),
					},
				))
			})

			It("parses IPv6 IP address", func() {
				ip, err := spiderpoolip.ParseIP(constant.IPv6, "abcd:1234::1", false)
				Expect(err).NotTo(HaveOccurred())
				Expect(ip).To(Equal(
					&net.IPNet{
						IP:   net.ParseIP("abcd:1234::1"),
						Mask: net.CIDRMask(128, 128),
					},
				))
			})
		})

		Describe("CIDR format", func() {
			When("Verifying", func() {
				It("inputs invalid IP version", func() {
					ip, err := spiderpoolip.ParseIP(invalidIPVersion, "172.18.40.40/24", true)
					Expect(err).To(MatchError(spiderpoolip.ErrInvalidIPVersion))
					Expect(ip).To(BeNil())
				})

				It("inputs invalid CIDR address", func() {
					ip, err := spiderpoolip.ParseIP(constant.IPv4, invalidCIDR, true)
					Expect(err).To(MatchError(spiderpoolip.ErrInvalidCIDRFormat))
					Expect(ip).To(BeNil())
				})
			})

			It("parses IPv4 CIDR address", func() {
				ip, err := spiderpoolip.ParseIP(constant.IPv4, "172.18.40.40/24", true)
				Expect(err).NotTo(HaveOccurred())
				Expect(ip).To(Equal(
					&net.IPNet{
						IP:   net.IPv4(172, 18, 40, 40),
						Mask: net.CIDRMask(24, 32),
					},
				))
			})

			It("parses IPv6 CIDR address", func() {
				ip, err := spiderpoolip.ParseIP(constant.IPv6, "abcd:1234::1/120", true)
				Expect(err).NotTo(HaveOccurred())
				Expect(ip).To(Equal(
					&net.IPNet{
						IP:   net.ParseIP("abcd:1234::1"),
						Mask: net.CIDRMask(120, 128),
					},
				))
			})
		})
	})

	Describe("Test ContainsIP", func() {
		When("Verifying", func() {
			It("inputs invalid IP version", func() {
				contains, err := spiderpoolip.ContainsIP(invalidIPVersion, "172.18.40.0/24", "172.18.40.40")
				Expect(err).To(MatchError(spiderpoolip.ErrInvalidIPVersion))
				Expect(contains).To(BeFalse())
			})

			It("inputs invalid subnet", func() {
				contains, err := spiderpoolip.ContainsIP(constant.IPv4, invalidCIDR, "172.18.40.40")
				Expect(err).To(MatchError(spiderpoolip.ErrInvalidCIDRFormat))
				Expect(contains).To(BeFalse())
			})

			It("inputs invalid IP address", func() {
				contains, err := spiderpoolip.ContainsIP(constant.IPv4, "172.18.40.0/24", invalidIP)
				Expect(err).To(MatchError(spiderpoolip.ErrInvalidIPFormat))
				Expect(contains).To(BeFalse())
			})
		})

		When("IPv4", func() {
			It("tests that a subnet contains the IP address", func() {
				contains, err := spiderpoolip.ContainsIP(constant.IPv4, "172.18.40.0/24", "172.18.40.40")
				Expect(err).NotTo(HaveOccurred())
				Expect(contains).To(BeTrue())
			})

			It("test that a subnet does not contain the IP address", func() {
				contains, err := spiderpoolip.ContainsIP(constant.IPv4, "172.18.41.0/24", "172.18.40.40")
				Expect(err).NotTo(HaveOccurred())
				Expect(contains).To(BeFalse())
			})
		})

		When("IPv6", func() {
			It("tests that a subnet contains the IP address", func() {
				contains, err := spiderpoolip.ContainsIP(constant.IPv6, "abcd:1234::/120", "abcd:1234::1")
				Expect(err).NotTo(HaveOccurred())
				Expect(contains).To(BeTrue())
			})

			It("test that a subnet does not contain the IP address", func() {
				contains, err := spiderpoolip.ContainsIP(constant.IPv6, "abcd:1235::/120", "abcd:1234::1")
				Expect(err).NotTo(HaveOccurred())
				Expect(contains).To(BeFalse())
			})
		})
	})

	Describe("Test IsIP", func() {
		When("Verifying", func() {
			It("inputs invalid IP version", func() {
				err := spiderpoolip.IsIP(invalidIPVersion, "172.18.40.40")
				Expect(err).To(MatchError(spiderpoolip.ErrInvalidIPVersion))
			})

			It("inputs invalid IP address", func() {
				err := spiderpoolip.IsIP(constant.IPv4, invalidIP)
				Expect(err).To(MatchError(spiderpoolip.ErrInvalidIPFormat))
			})
		})

		It("is IPv4 IP address", func() {
			Expect(spiderpoolip.IsIP(constant.IPv4, "172.18.40.40")).To(Succeed())
		})

		It("is IPv6 IP address", func() {
			Expect(spiderpoolip.IsIP(constant.IPv6, "abcd:1234::1")).To(Succeed())
		})
	})

	Describe("Test IPsDiffSet", func() {
		It("finds the difference set of two groups of IPv4 addresses", func() {
			ips := spiderpoolip.IPsDiffSet(
				[]net.IP{
					net.IPv4(172, 18, 40, 1),
					net.IPv4(172, 18, 40, 2),
				},
				[]net.IP{
					net.IPv4(172, 18, 40, 2),
					net.IPv4(172, 18, 40, 3),
				},
			)
			Expect(ips).To(Equal([]net.IP{net.IPv4(172, 18, 40, 1)}))
		})

		It("finds the difference set of two groups of IPv6 addresses", func() {
			ips := spiderpoolip.IPsDiffSet(
				[]net.IP{
					net.ParseIP("abcd:1234::1"),
					net.ParseIP("abcd:1234::2"),
				},
				[]net.IP{
					net.ParseIP("abcd:1234::2"),
					net.ParseIP("abcd:1234::3"),
				},
			)
			Expect(ips).To(Equal([]net.IP{net.ParseIP("abcd:1234::1")}))
		})
	})

	Describe("Test IPsUnionSet", func() {
		It("finds the union set of two groups of IPv4 addresses", func() {
			ips := spiderpoolip.IPsUnionSet(
				[]net.IP{
					net.IPv4(172, 18, 40, 1),
					net.IPv4(172, 18, 40, 2),
				},
				[]net.IP{
					net.IPv4(172, 18, 40, 2),
					net.IPv4(172, 18, 40, 3),
				},
			)
			Expect(ips).To(Equal(
				[]net.IP{
					net.IPv4(172, 18, 40, 1),
					net.IPv4(172, 18, 40, 2),
					net.IPv4(172, 18, 40, 3),
				}),
			)
		})

		It("finds the union set of two groups of IPv6 addresses", func() {
			ips := spiderpoolip.IPsUnionSet(
				[]net.IP{
					net.ParseIP("abcd:1234::1"),
					net.ParseIP("abcd:1234::2"),
				},
				[]net.IP{
					net.ParseIP("abcd:1234::2"),
					net.ParseIP("abcd:1234::3"),
				},
			)
			Expect(ips).To(Equal(
				[]net.IP{
					net.ParseIP("abcd:1234::1"),
					net.ParseIP("abcd:1234::2"),
					net.ParseIP("abcd:1234::3"),
				}),
			)
		})
	})

	Describe("Test IPsIntersectionSet", func() {
		It("finds the intersection set of two groups of IPv4 addresses", func() {
			ips := spiderpoolip.IPsIntersectionSet(
				[]net.IP{
					net.IPv4(172, 18, 40, 1),
					net.IPv4(172, 18, 40, 2),
				},
				[]net.IP{
					net.IPv4(172, 18, 40, 2),
					net.IPv4(172, 18, 40, 3),
				},
			)
			Expect(ips).To(Equal([]net.IP{net.IPv4(172, 18, 40, 2)}))
		})

		It("finds the intersection set of two groups of IPv6 addresses", func() {
			ips := spiderpoolip.IPsIntersectionSet(
				[]net.IP{
					net.ParseIP("abcd:1234::1"),
					net.ParseIP("abcd:1234::2"),
				},
				[]net.IP{
					net.ParseIP("abcd:1234::2"),
					net.ParseIP("abcd:1234::3"),
				},
			)
			Expect(ips).To(Equal([]net.IP{net.ParseIP("abcd:1234::2")}))
		})
	})

	Describe("Test NextIP", func() {
		It("returns the next IP address of the IPv4 IP address", func() {
			ip := spiderpoolip.NextIP(net.IPv4(172, 18, 40, 40))
			Expect(ip).To(Equal(net.IPv4(172, 18, 40, 41)))
		})

		It("returns the next IP address of the IPv6 IP address", func() {
			ip := spiderpoolip.NextIP(net.ParseIP("abcd:1234::1"))
			Expect(ip).To(Equal(net.ParseIP("abcd:1234::2")))
		})
	})

	Describe("Test PrevIP", func() {
		It("returns the previous IP address of the IPv4 IP address", func() {
			ip := spiderpoolip.PrevIP(net.IPv4(172, 18, 40, 40))
			Expect(ip).To(Equal(net.IPv4(172, 18, 40, 39)))
		})

		It("returns the previous IP address of the IPv6 IP address", func() {
			ip := spiderpoolip.PrevIP(net.ParseIP("abcd:1234::1"))
			Expect(ip).To(Equal(net.ParseIP("abcd:1234::0")))
		})
	})

	Describe("Test Cmp", func() {
		It("compares two IPv4 IP addresses", func() {
			Expect(spiderpoolip.Cmp(
				net.IPv4(172, 18, 40, 1),
				net.IPv4(172, 18, 40, 2),
			)).To(BeNumerically("<", 0))

			Expect(spiderpoolip.Cmp(
				net.IPv4(172, 18, 40, 1),
				net.IPv4(172, 18, 40, 1),
			)).To(BeNumerically("==", 0))

			Expect(spiderpoolip.Cmp(
				net.IPv4(172, 18, 40, 2),
				net.IPv4(172, 18, 40, 1),
			)).To(BeNumerically(">", 0))
		})

		It("compares two IPv6 IP addresses", func() {
			Expect(spiderpoolip.Cmp(
				net.ParseIP("abcd:1234::1"),
				net.ParseIP("abcd:1234::2"),
			)).To(BeNumerically("<", 0))

			Expect(spiderpoolip.Cmp(
				net.ParseIP("abcd:1234::1"),
				net.ParseIP("abcd:1234::1"),
			)).To(BeNumerically("==", 0))

			Expect(spiderpoolip.Cmp(
				net.ParseIP("abcd:1234::2"),
				net.ParseIP("abcd:1234::1"),
			)).To(BeNumerically(">", 0))
		})
	})
})
