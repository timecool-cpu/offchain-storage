package main

type Graph struct {
	pk    []byte
	index int64
	log2  int64
	pow2  int64
	size  int64
}

func (g *Graph) GetParents(node, index int64) []int64 {
	if node < int64(1<<uint64(index)) {
		return nil
	}

	offset0, offset1 := g.GetGraph(node, index)

	var res []int64
	if offset0 != 0 {
		res = append(res, node-offset0)
	}
	if offset1 != 0 {
		res = append(res, node-offset1)
	}
	return res
}

// get graph that node belongs to, so i can find the parents
func (g *Graph) GetGraph(node, index int64) (int64, int64) {
	if index == 1 {
		if node < 2 {
			return 2, 0
		} else if node == 2 {
			return 1, 2
		} else if node == 3 {
			return 3, 2
		}
	}

	pow2index := int64(1 << uint64(index))
	pow2index_1 := int64(1 << uint64(index-1))
	sources := pow2index
	firstButter := sources + numButterfly(index-1)
	firstXi := firstButter + numXi(index-1)
	secondXi := firstXi + numXi(index-1)
	secondButter := secondXi + numButterfly(index-1)
	sinks := secondButter + sources

	if node < sources {
		return pow2index, 0
	} else if node >= sources && node < firstButter {
		if node < sources+pow2index_1 {
			return pow2index, pow2index_1
		} else {
			parent0, parent1 := g.ButterflyParents(sources, node, index)
			return node - parent0, node - parent1
		}
	} else if node >= firstButter && node < firstXi {
		node = node - firstButter
		return g.GetGraph(node, index-1)
	} else if node >= firstXi && node < secondXi {
		node = node - firstXi
		return g.GetGraph(node, index-1)
	} else if node >= secondXi && node < secondButter {
		if node < secondXi+pow2index_1 {
			return pow2index_1, 0
		} else {
			parent0, parent1 := g.ButterflyParents(secondXi, node, index)
			return node - parent0, node - parent1
		}
	} else if node >= secondButter && node < sinks {
		offset := (node - secondButter) % pow2index_1
		parent1 := sinks - numXi(index) + offset
		if offset+secondButter == node {
			return pow2index_1, node - parent1
		} else {
			return pow2index, node - parent1 - pow2index_1
		}
	} else {
		return 0, 0
	}
}

// compute the offsets for the two parents in the butterfly graph
func (g *Graph) ButterflyParents(begin, node, index int64) (int64, int64) {
	pow2index_1 := int64(1 << uint64(index-1))
	level := (node - begin) / pow2index_1
	var prev int64
	shift := (index - 1) - level
	if level > (index - 1) {
		shift = level - (index - 1)
	}
	i := (node - begin) % pow2index_1
	if (i>>uint64(shift))&1 == 0 {
		prev = i + (1 << uint64(shift))
	} else {
		prev = i - (1 << uint64(shift))
	}
	parent0 := begin + (level-1)*pow2index_1 + prev
	parent1 := node - pow2index_1
	return parent0, parent1
}

func numButterfly(index int64) int64 {
	return 2 * (1 << uint64(index)) * index
}

func numXi(index int64) int64 {
	return (1 << uint64(index)) * (index + 1) * index
}
