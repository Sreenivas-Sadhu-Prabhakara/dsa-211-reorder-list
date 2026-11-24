# Reorder List

> NeetCode 150 · Linked List · Blind 75 — three known moves composed: middle via slow/fast, reverse the back half, interleave. Go, stdlib only, tested.

## Problem

Rearrange L0→L1→…→Ln into L0→Ln→L1→Ln−1→… in place — values untouched, only pointers.

## Approach

No new algorithm, which is the point of its Blind-75 slot: it tests whether three primitives compose cleanly. Slow/fast pointers find the end of the first half (the `fast.Next != nil && fast.Next.Next != nil` guard puts the extra middle node of odd lists into the *first* half); cut the list there; reverse the second half (dsa-025's loop verbatim); then zip the halves together, always saving both next-pointers before rewiring. The interleave terminates on the reversed half running out, which the split guaranteed is never the longer one. Palindrome-list (dsa-034) is this same pipeline with a comparison instead of a zip.

**Complexity:** O(n) time, O(1) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
