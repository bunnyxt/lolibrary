<?php

namespace App\Http\Controllers\Items;

use App\Models\Tag;
use App\Http\Controllers\Controller;

class TagController extends Controller
{
    /**
     * Show a tag.
     *
     * @param \App\Models\Tag $tag
     * @return \Illuminate\Http\Response
     */
    public function show(Tag $tag)
    {
        return view('tags.show', compact('tag'));
    }

    /**
     * Show a paginated list of tags.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        // todo: make this a static ::index() method.
        $tags = Tag::paginate(36);

        return view('tags.index', compact('tags'));
    }
}
